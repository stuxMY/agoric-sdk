// @ts-check

import { details as X } from '@agoric/assert';
import { E } from '@agoric/eventual-send';
import { Far } from '@agoric/marshal';
import { makePromiseKit } from '@agoric/promise-kit';
import { sameStructure } from '@agoric/same-structure';

import {
  ChoiceMethod,
  QuorumRule,
  ElectionType,
  looksLikeBallotSpec,
} from './ballotBuilder.js';
import { looksLikeParam } from './paramManager.js';

/** @type {MakeParamChangePositions} */
const makeParamChangePositions = (paramSpec, proposedValue) => {
  const positive = harden({ changeParam: paramSpec, proposedValue });
  const negative = harden({ noChange: paramSpec });
  return { positive, negative };
};

/** @type {ValidateParamChangeBallot} */
const validateParamChangeBallot = details => {
  assert(
    details.method === ChoiceMethod.CHOOSE_N,
    X`ChoiceMethod must be CHOOSE_N, not ${details.method}`,
  );
  assert(
    details.electionType === ElectionType.PARAM_CHANGE,
    X`ElectionType must be PARAM_CHANGE, not ${details.electionType}`,
  );
  assert(
    details.maxChoices === 1,
    X`maxChoices must be 1, not ${details.maxChoices}`,
  );
  assert(
    details.quorumRule === QuorumRule.MAJORITY,
    X`QuorumRule must be MAJORITY, not ${details.quorumRule}`,
  );
  assert(
    details.tieOutcome.noChange,
    X`tieOutcome must be noChange, not ${details.tieOutcome}`,
  );
};

/** @type {AssertBallotConcernsQuestion} */
const assertBallotConcernsQuestion = (paramName, ballotDetails) => {
  assert(
    // @ts-ignore typescript isn't sure the question is a paramChangeQuestion
    // if it isn't, the assertion will fail.
    ballotDetails.question.paramSpec.parameterName === paramName,
    X`expected ${paramName} to be included`,
  );
};

/** @type {SetupGovernance} */
const setupGovernance = async (
  paramManagerAccessor,
  poserFacet,
  contractInstance,
  timer,
) => {
  /** @type {WeakSet<Instance>} */
  const ballotCounters = new WeakSet();

  /** @type {VoteOnParamChange} */
  const voteOnParamChange = async (
    paramSpec,
    proposedValue,
    ballotCounterInstallation,
    deadline,
  ) => {
    const paramMgr = E(paramManagerAccessor).get(paramSpec);
    const paramName = paramSpec.parameterName;
    looksLikeParam(proposedValue, paramName);
    const outcomeOfUpdateP = makePromiseKit();

    const { positive, negative } = makeParamChangePositions(
      paramSpec,
      proposedValue,
    );
    const question = harden({
      paramSpec,
      contract: contractInstance,
      proposedValue,
    });
    const ballotSpec = looksLikeBallotSpec({
      method: ChoiceMethod.CHOOSE_N,
      question,
      positions: [positive, negative],
      electionType: ElectionType.PARAM_CHANGE,
      maxChoices: 1,
      closingRule: { timer, deadline },
      quorumRule: QuorumRule.MAJORITY,
      tieOutcome: negative,
    });

    const {
      publicFacet: counterPublicFacet,
      instance: ballotCounter,
    } = await E(poserFacet).addQuestion(ballotCounterInstallation, ballotSpec);

    ballotCounters.add(ballotCounter);

    E(counterPublicFacet)
      .getOutcome()
      .then(outcome => {
        if (sameStructure(positive, outcome)) {
          E(paramMgr)
            [`update${paramName}`](proposedValue)
            .then(newValue => outcomeOfUpdateP.resolve(newValue))
            .catch(e => {
              outcomeOfUpdateP.reject(e);
            });
        }
      })
      .catch(e => {
        outcomeOfUpdateP.reject(e);
      });

    return {
      outcomeOfUpdate: outcomeOfUpdateP.promise,
      instance: ballotCounter,
      details: E(counterPublicFacet).getDetails(),
    };
  };

  return Far('paramGovernor', {
    voteOnParamChange,
    createdBallot: b => ballotCounters.has(b),
  });
};

harden(setupGovernance);
harden(makeParamChangePositions);
harden(validateParamChangeBallot);
harden(assertBallotConcernsQuestion);
export {
  setupGovernance,
  makeParamChangePositions,
  validateParamChangeBallot,
  assertBallotConcernsQuestion,
};