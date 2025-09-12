package query

var GetBrandListData = `SELECT
  ba."refApplicationId",
  ba."refBrandId",
  ba."refApplicationCustId",
  ba."refBrandName",
  bc."refBrandCategoryName",
  bcd."refBrandContactPerson",
  wh."refIfWareHouse",
  ba."refCreateAT",
  aps."refApplicationStatus",
  bd."refLogo",
  ba."refUpdateAt"
FROM
  brand."refBrandApplication" ba
  LEFT JOIN brand."refSocialMedia" sm ON CAST(sm."refSocialMediaId" AS INTEGER) = ba."refSocialMediaId"
  LEFT JOIN brand."refBrandCategory" bc ON CAST(bc."refBrandCategoryId" AS INTEGER) = ba."refProductCatageoryId"
  LEFT JOIN brand."refBrandContactDetails" bcd ON CAST(bcd."refBrandContactId" AS INTEGER) = ba."refBrandLocationId"
  LEFT JOIN brand."refBrandLocation" bl ON CAST(bl."refLocationId" AS INTEGER) = ba."refBrandLocationId"
  LEFT JOIN brand."refBrandWareHouse" wh ON CAST(wh."refWareHouseId" AS INTEGER) = ba."refWareHoueId"
  LEFT JOIN brand."refDocuments" bd ON CAST(bd."refDocumentsId" AS INTEGER) = ba."refDocumentsId"
  LEFT JOIN brand."refApplicationStatus" aps ON CAST(aps."refApplicationStatusId" AS INTEGER) = ba."refApplicationStatus"
WHERE
  ba."refApplicationStatus" = $1
  AND ba."refProductCatageoryId" = ANY($2)`

var GetBrandDataFromDbQuery = `SELECT
  ba."refBrandName",
  ba."refProductCatageoryId",
  bd."refLogo",
  ba."refBrandDesciption",
  bs."refWebsiteUrl",
  bs."refInstaUrl",
  bc."refBrandContactPerson",
  bc."refDesignation",
  bc."refBrandMobile",
  bc."refBrandEmail",
  bl."refAddress",
  bl."refCity",
  bl."refZipCode",
  bl."refState",
  bd."refAddressProf",
  ba."refGstin",
  ba."refCin",
  bd."refGstin" AS "refGstinPath",
  bd."refPanCars" AS "refPanCars",
  bw."refIfWareHouse" AS "refIfWareHouse",
  bw."refAddress" AS "wareHouseAddress",
  bw."refCity" AS "wareHouseCity",
  bw."refDistrict" AS "wareHouseDistrict",
  bw."refZipcode" AS "wareHouseZipCode",
  bw."refState" AS "wareHouseState",
  ba."refSaveDraft",
  aps."refApplicationStatusId",
  aps."refApplicationStatus",
  bd."refAddressProfName",
  bd."refPanCarsName",
  bd."refGstinName",
  bd."refLogoName"
FROM
  brand."refBrandApplication" ba
  LEFT JOIN brand."refBrandContactDetails" bc ON CAST(ba."refBrandContactId" AS INTEGER) = bc."refBrandContactId"
  LEFT JOIN brand."refBrandLocation" bl ON CAST(bl."refLocationId" AS INTEGER) = ba."refBrandLocationId"
  LEFT JOIN brand."refDocuments" bd ON CAST(bd."refDocumentsId" AS INTEGER) = ba."refDocumentsId"
  LEFT JOIN brand."refBrandWareHouse" bw ON CAST(bw."refWareHouseId" AS INTEGER) = ba."refWareHoueId"
  LEFT JOIN brand."refSocialMedia" bs ON CAST(bs."refSocialMediaId" AS INTEGER) = ba."refSocialMediaId"
  LEFT JOIN brand."refApplicationStatus" aps ON CAST(aps."refApplicationStatusId" AS INTEGER) = ba."refApplicationStatus"
WHERE
  ba."refApplicationId" = $1`

var UpdateBrandApprovalStatusQuery = `UPDATE brand."refBrandApplication" ba
SET
  "refApplicationStatus" = $2,
  "refApplicationStatusDesciption" = $3,
  "refBrandPassword" = $4,
  "refBrandHashedPassword" = $5,
  "refUpdateAt" = $6
FROM brand."refBrandContactDetails" bc
WHERE ba."refApplicationId" = $1
  AND CAST(bc."refBrandContactId" AS INTEGER) = ba."refBrandContactId"
RETURNING
  ba."refApplicationCustId",
  ba."refBrandName",
  bc."refBrandEmail",
  bc."refBrandContactPerson",
  ba."refApplicationCustId";
`

var UpdateBrandRejectedStatus = `UPDATE
  brand."refBrandApplication" ba
SET
  "refApplicationStatus" = $2,
  "refApplicationStatusDesciption" = $3,
  "refUpdateAt" = $4
FROM
  brand."refBrandContactDetails" bc
WHERE
  ba."refApplicationId" = $1
  AND CAST(bc."refBrandContactId" AS INTEGER) = ba."refBrandContactId"
RETURNING
  ba."refApplicationCustId",
  ba."refBrandName",
  bc."refBrandEmail",
  bc."refBrandContactPerson";`
