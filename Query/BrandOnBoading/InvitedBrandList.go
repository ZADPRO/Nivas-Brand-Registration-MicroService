package brandRegistrationQuery

var InvitedBrandListQuery = `SELECT
  ba."refApplicationCustId",
  ba."refApplicationId",
  ba."refBrandName",
  ba."refMailId",
  TO_CHAR(ba."refCreateAT"::DATE, 'Mon DD, YYYY') AS "refCreateAT",
  bc."refBrandContactPerson",
  CASE
    WHEN ba."refApplicationStatus" = 1 THEN 'Pending'
    ELSE 'Accept'
  END AS "refStatus"
FROM
  brand."refBrandApplication" ba
  LEFT JOIN brand."refBrandContactDetails" bc ON CAST(bc."refBrandContactId" AS INTEGER) = ba."refBrandContactId"
WHERE
  ba."refApplicationStatus" IN (1, 2, 6);`

var ReInviteBrandDataQuery = `SELECT
  ba."refApplicationCustId",
  ba."refApplicationId",
  ba."refBrandName",
  ba."refMailId",
  TO_CHAR(ba."refCreateAT"::DATE, 'Mon DD, YYYY') AS "refCreateAT",
  bc."refBrandContactPerson",
  CASE
    WHEN ba."refApplicationStatus" = 1 THEN 'Pending'
    ELSE 'Accept'
  END AS "refStatus"
FROM
  brand."refBrandApplication" ba
  LEFT JOIN brand."refBrandContactDetails" bc ON CAST(bc."refBrandContactId" AS INTEGER) = ba."refBrandContactId"
WHERE
  ba."refApplicationId" = $1`
