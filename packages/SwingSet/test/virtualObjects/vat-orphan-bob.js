import { Far } from '@endo/marshal';
import { defineKind } from '@agoric/vat-data';

export function buildRootObject(vatPowers) {
  const { testLog } = vatPowers;

  let extracted;

  const makeThing = defineKind('thing', () => ({}), {
    workingFacet: {
      statelessMethod: () => 0,
      extractState: ({ state }) => {
        extracted = state;
      },
      extractCohort: ({ facets }) => {
        extracted = facets;
      },
    },
    unusedFacet: {
      method: () => 0,
    },
  });

  let strongRetainer;
  const weakRetainer = new WeakSet();
  let retentionMode;

  return Far('root', {
    retain(mode) {
      retentionMode = mode;
      const originalFacet = makeThing().workingFacet;
      switch (mode) {
        case 'facet':
          strongRetainer = originalFacet;
          break;
        case 'wfacet':
          weakRetainer.add(originalFacet);
          break;
        case 'method':
          strongRetainer = originalFacet.statelessMethod;
          break;
        case 'wmethod':
          weakRetainer.add(originalFacet.statelessMethod);
          break;
        case 'proto':
          // eslint-disable-next-line no-proto
          strongRetainer = originalFacet.__proto__;
          break;
        case 'wproto':
          // eslint-disable-next-line no-proto
          weakRetainer.add(originalFacet.__proto__);
          break;
        case 'cohort':
          originalFacet.extractCohort();
          strongRetainer = extracted;
          extracted = null;
          break;
        case 'wcohort':
          originalFacet.extractCohort();
          weakRetainer.add(extracted);
          extracted = null;
          break;
        case 'state':
          originalFacet.extractState();
          strongRetainer = extracted;
          extracted = null;
          break;
        case 'wstate':
          originalFacet.extractState();
          weakRetainer.add(extracted);
          extracted = null;
          break;
        default:
          break;
      }
      makeThing(); // push original out of the cache
      return originalFacet;
    },
    testForRetention(thing) {
      let compare;
      switch (retentionMode) {
        case 'facet':
          compare = strongRetainer === thing;
          break;
        case 'wfacet':
          compare = weakRetainer.has(thing);
          break;
        case 'method':
          compare = strongRetainer === thing.statelessMethod;
          break;
        case 'wmethod':
          compare = weakRetainer.has(thing.statelessMethod);
          break;
        case 'proto':
          // eslint-disable-next-line no-proto
          compare = strongRetainer === thing.__proto__;
          break;
        case 'wproto':
          // eslint-disable-next-line no-proto
          compare = weakRetainer.has(thing.__proto__);
          break;
        case 'cohort':
          thing.extractCohort();
          compare = strongRetainer === extracted;
          break;
        case 'wcohort':
          thing.extractCohort();
          compare = weakRetainer.has(extracted);
          break;
        case 'state':
          thing.extractState();
          compare = strongRetainer === extracted;
          break;
        case 'wstate':
          thing.extractState();
          compare = weakRetainer.has(extracted);
          break;
        default:
          break;
      }
      testLog(`compare old === new : ${compare}`);
    },
  });
}
