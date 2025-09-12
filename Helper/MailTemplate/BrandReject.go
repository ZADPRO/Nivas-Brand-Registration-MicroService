package BrandInviteMailTemplate

var BrandRejectTemplate = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@400;700&display=swap" rel="stylesheet">

    <title>Nivas - Rejection</title>
    <style>
      body {
        margin: 0;
        padding: 0;
        font-family: Arial;
        background-color: #f5f5f5;
        width:600px;
      }
      .container2 {
        width: 95%;
        margin: 20px 20px 15px 15px;
        background: #ffffff;
        border-radius: 8px;
        overflow: hidden;
        border: 1;
        border: #f0eeec;
        text-align: justify
      }
      .header {
        text-align: start;
        padding: 20px;
      }
      .header img.logo {
        width: 105px;
        height: 32;
        opacity: 1;
        margin: 10px 10px 20px 0px ;
      }
      .header h2 {
        font-family: 'Montserrat', sans-serif;
        font-weight: 700;
        font-size: 29px;
        line-height: 36px;
        letter-spacing: -1px;
        text-align: center;
        margin: 0;
      }
      .header p {
        color: #6c6c6c;
        font-size: 15px;
        text-align: center;
        font-weight: 400;
      }
      .hero {
        display: flex;
        justify-content: center;
      }
      .hero img {
        width: 100%;
      }
      .content {
        padding: 20px;
        font-size: 17px;
        color: #27282c;
        line-height: 26px;
        letter-spacing: 1%;
        font-weight: 400;
      }
      .btn {
        display: inline-block;
        background: #000000;
        color: #fff !important;
        text-decoration: none;
        padding: 12px 24px;
        border-radius: 32px;
        margin-top: 20px;
        font-size: 14px;
      }
      .footer {
        padding: 2px 15px 10px 15px;
        background: #fafafa;
        font-size: 17px;
        text-align: left;
        color: #444;
      }
      .footer strong {
        font-weight: 700;
        font-size: 17px;
        line-height: 26px;
        letter-spacing: 1%;
      }
      .footer img {
        width: 105px;
        height: 32;
        opacity: 1;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <!-- Header -->
      <div class="header">
        <img src="https://raw.githubusercontent.com/vijayzadroit/nivas/refs/heads/main/logo.png" alt="nivas logo" class="logo" />
        <h2>Resubmission Needed to Complete Your Onboarding.</h2>
      </div>

      <!-- Hero Image -->
      <div class="hero">
        <img src="https://raw.githubusercontent.com/vijayzadroit/nivas/main/RejectedHero.png" alt="Banner" />
      </div>

      <!-- Content -->
      <div class="container2">
        <div class="content">
          <span style="font-size: 18px; font-weight: 700">
            Dear {{.ContactPerson}},
          </span>
          <p>
            Thank you for your interest in partnering with
            <strong>Nivas - House of Celebrities</strong> and for submitting your
            application to register <strong>{{.BrandName}}</strong> on our
            platform. We truly appreciate the time and effort you invested in
            your submission.
          </p>
          <p>
            After careful review of your application, we have found that there
            are errors in your submission.<br />
            <strong> Rejection Reason - {{.RejectionReason}} <br /> </strong>
            We request that you correct the details and re-submit the
            application.
          </p>

          <p>
            Thank you again for considering
            <strong>Nivas - House of Celebrities</strong> as a platform for
            your brand. Please reach out to us if you have further questions or
            concerns.
          </p>

          <a href="{{.ResubmitURL}}" class="btn">Edit and Resubmit application</a>
        </div>

        <!-- Footer -->
        <div class="footer">
          <p>Warm regards,</p>
          <p>
            <img
              src="https://raw.githubusercontent.com/vijayzadroit/nivas/refs/heads/main/logo.png"
              alt="nivas logo"
              class="logo"
              style="width: 107px; height: 30px; padding-right: 10px; transform: rotate(0deg); opacity: 1;"
            />
            <br />
            <strong>
              Nivas - House of Celebrities<br />
              {{.BrandMailId}}<br />
              {{.BrandMobile}}<br />
              <a href="https://www.nivasinc.com/">https://www.nivasinc.com/</a>
            </strong>
          </p>
        </div>
      </div>
    </div>
  </body>
</html>`
