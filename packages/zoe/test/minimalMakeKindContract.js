/* global VatData */

const start = _zcf => {
  VatData.defineKind();
  const kh = VatData.makeKindHandle();
  VatData.defineDurableKind(kh);
  VatData.makeScalarBigMapStore();
  VatData.makeScalarBigWeakMapStore();
  VatData.makeScalarBigSetStore();
  VatData.makeScalarBigWeakSetStore();

  return harden({});
};
harden(start);
export { start };
