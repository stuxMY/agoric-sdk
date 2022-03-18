// @ts-check

import '@agoric/zoe/src/types.js';

import { makeIssuerKit, AssetKind, AmountMath } from '@agoric/ertp';
import { E } from '@endo/far';

import { assert } from '@agoric/assert';
import buildManualTimer from '@agoric/zoe/tools/manualTimer.js';
import { makeFakePriceAuthority } from '@agoric/zoe/tools/fakePriceAuthority.js';
import {
  makeRatio,
  makeRatioFromAmounts,
  multiplyRatios,
  floorDivideBy,
  makeUnitAmount,
} from '@agoric/zoe/src/contractSupport/index.js';
import { Far } from '@endo/marshal';

import { makeNotifierKit } from '@agoric/notifier';
import { makeInnerVault } from '../../src/vaultFactory/vault.js';
import { paymentFromZCFMint } from '../../src/vaultFactory/burn.js';

const BASIS_POINTS = 10000n;
const SECONDS_PER_HOUR = 60n * 60n;
const DAY = SECONDS_PER_HOUR * 24n;

/**
 * @param {ZCF} zcf
 * @param {{feeMintAccess: FeeMintAccess}} privateArgs
 */
export async function start(zcf, privateArgs) {
  console.log(`contract started`);
  assert.typeof(privateArgs.feeMintAccess, 'object');

  const collateralKit = makeIssuerKit('Collateral');
  const { brand: collateralBrand } = collateralKit;
  await zcf.saveIssuer(collateralKit.issuer, 'Collateral'); // todo: CollateralETH, etc

  const runMint = await zcf.registerFeeMint('RUN', privateArgs.feeMintAccess);
  const { brand: runBrand } = runMint.getIssuerRecord();

  const runDecimalPlaces = await E.get(E(runBrand).getDisplayInfo())
    .decimalPlaces;
  const { brand: priceBrand } = makeIssuerKit(
    'USD',
    AssetKind.NAT,
    harden({
      decimalPlaces: (runDecimalPlaces || 0) + 3, // Arbitrarily different than runDecimalPlaces.
    }),
  );

  const debtPriceRatio = await Promise.all(
    [runBrand, priceBrand].map(makeUnitAmount),
  ).then(async ([debt, price]) => makeRatioFromAmounts(debt, price));
  const { zcfSeat: vaultFactorySeat } = zcf.makeEmptySeatKit();

  let vaultCounter = 0;

  let currentInterest = makeRatio(5n, runBrand); // 5%
  let compoundedInterest = makeRatio(100n, runBrand); // starts at 1.0, no interest

  function reallocateWithFee(amount, fromSeat, otherSeat) {
    vaultFactorySeat.incrementBy(
      fromSeat.decrementBy(
        harden({
          RUN: amount,
        }),
      ),
    );
    if (otherSeat !== undefined) {
      zcf.reallocate(vaultFactorySeat, fromSeat, otherSeat);
    } else {
      zcf.reallocate(vaultFactorySeat, fromSeat);
    }
  }

  /** @type {Parameters<typeof makeInnerVault>[1]} */
  const managerMock = Far('vault manager mock', {
    getLiquidationMargin() {
      return makeRatio(105n, runBrand);
    },
    getLoanFee() {
      return makeRatio(500n, runBrand, BASIS_POINTS);
    },
    getInterestRate() {
      return currentInterest;
    },
    getCollateralBrand() {
      return collateralBrand;
    },
    getChargingPeriod() {
      return DAY;
    },
    getRecordingPeriod() {
      return DAY;
    },
    reallocateWithFee,
    applyDebtDelta() {},
    getCollateralQuote() {
      return Promise.resolve({
        quoteAmount: AmountMath.make(runBrand, 0n),
        quotePayment: null,
      });
    },
    getCompoundedInterest: () => compoundedInterest,
    updateVaultPriority: () => {
      // noop
    },
  });

  const timer = buildManualTimer(console.log, 0n, DAY);
  const options = {
    actualBrandIn: collateralBrand,
    actualBrandOut: priceBrand,
    // Scale the priceList by the debtPriceRatio.
    priceList: [80].map(
      n =>
        floorDivideBy(AmountMath.make(runBrand, BigInt(n)), debtPriceRatio)
          .value,
    ),
    tradeList: undefined,
    timer,
    quoteMint: makeIssuerKit('quote', AssetKind.SET).mint,
  };
  const priceAuthority = makeFakePriceAuthority(options);

  const { notifier: managerNotifier } = makeNotifierKit();

  const innerVault = await makeInnerVault(
    zcf,
    managerMock,
    managerNotifier,
    // eslint-disable-next-line no-plusplus
    String(vaultCounter++),
    runMint,
    priceAuthority,
    debtPriceRatio,
  );

  const advanceRecordingPeriod = () => {
    timer.tick();

    // skip the debt calculation for this mock manager
    const currentInterestAsMultiplicand = makeRatio(
      100n + currentInterest.numerator.value,
      currentInterest.numerator.brand,
    );
    compoundedInterest = multiplyRatios(
      compoundedInterest,
      currentInterestAsMultiplicand,
    );
  };

  const setInterestRate = percent => {
    currentInterest = makeRatio(percent, runBrand);
  };

  zcf.setTestJig(() => ({
    advanceRecordingPeriod,
    collateralKit,
    runMint,
    setInterestRate,
    vault: innerVault,
  }));

  async function makeHook(seat) {
    const vaultKit = await innerVault.initVaultKit(seat);
    return {
      vault: innerVault,
      runMint,
      collateralKit,
      actions: Far('vault actions', {
        add() {
          return vaultKit.invitationMakers.AdjustBalances();
        },
      }),
      notifier: vaultKit.vaultNotifier,
    };
  }

  console.log(`makeContract returning`);

  const vaultAPI = Far('vaultAPI', {
    makeAdjustBalancesInvitation() {
      return innerVault.makeAdjustBalancesInvitation();
    },
    mintRun(amount) {
      return paymentFromZCFMint(zcf, runMint, amount);
    },
  });

  const testInvitation = zcf.makeInvitation(makeHook, 'foo');
  return harden({ creatorInvitation: testInvitation, creatorFacet: vaultAPI });
}
