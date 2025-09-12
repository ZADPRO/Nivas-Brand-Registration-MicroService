package BrandInviteMailTemplate

var BrandApproveMailTemplate = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@400;700&display=swap" rel="stylesheet">

    <title>Nivas - Invitation</title>
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
        margin: 15px 15px 15px 15px;
        background: #ffffff;
        border-radius: 8px;
        overflow: hidden;
        border: 1;
        border: #f0eeec;
        justify-content: center;
        /* box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1); */
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
        justify-content: center;
      }
      .content h3 {
        font-size: 16px;
        margin-bottom: 10px;
        color: #222;
      }
      .content ul {
        padding-left: 18px;
      }
      .content ul li {
        margin-bottom: 8px;
      }
      .btn {
        width: 218;
        height: 40;
        padding-top: 8px;
        padding-right: 24px;
        padding-bottom: 8px;
        padding-left: 24px;

        display: inline-block;
        background: #000000;
        color: #fff !important;
        text-decoration: none;
        /* padding: 12px 24px; */
        border-radius: 32px;
        margin-top: 20px;
        font-size: 14px;
      }
      .footer {
        padding: 2px 15px 10px 15px;
        font-size: 17px;
        text-align: left;
        color: #444;
      }
      .footer strong {
        font-weight: 700;
        font-style: Bold;
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
        <h2>Congratulations! Your Brand is Now Officially Part of Nivas</h2>
      </div>

      <!-- Hero Image -->
      <div class="hero">
        <img src="https://raw.githubusercontent.com/vijayzadroit/nivas/refs/heads/main/ApprovedHero.png" alt="Banner" />
      </div>

      <!-- Content -->
      <div class="container2">
        <div class="content">
          <span style="font-size: 18px; font-weight: 700">
            Dear {{.ContactPerson}},
          </span>
          <p>
            Thank you for choosing
            <strong>Nivas- House of Celebrities </strong>as a platform to
            showcase and grow <strong>{{.BrandName}}.</strong>   We are delighted
            to inform you that your registration has been approved.
          </p>
          <p>
            Your brand will now be part of our curated marketplace, featuring
            exclusive celebrity-owned products, and will benefit from our
            marketing, customer engagement, and sales support initiatives.
          </p>
          <strong>Next Steps:</strong>
          <ol>
            <li>
              Please find the login credentials for the seller dashboard login
              <br />
              <strong>Username : {{.UserName}}</strong> <br />
              <strong>One Time Password : {{.Password}}</strong>
            </li>
            <li>
              Our onboarding team will reach out to guide you through
              <strong> product listing, pricing, and brand guidelines.</strong>
            </li>
            <li>
              Please ensure all brand assets (logos, imagery, and product
              details) meet our quality requirements to maintain a consistent
              customer experience.
            </li>
          </ol>

          <p>
            We are excited to partner with you and bring your products to a
            passionate and engaged audience. If you have any questions or
            require assistance during onboarding, please reach out to your
            dedicated brand manager at  <strong> {{.BrandMailId}}.</strong>
          </p>

          <p>
            Thank you again for trusting <strong>
              Nivas - House of Celebrities.
            </strong>
            We look forward to a successful and long-lasting collaboration.
          </p>
          <a href="{{.WebSiteUrl}}" class="btn">Sign Up to the platform</a>
        </div>

        <!-- Footer -->
        <div class="footer">
          <p>Warm regards,</p>
          <p>
            <img
              src="https://raw.githubusercontent.com/vijayzadroit/nivas/refs/heads/main/logo.png"
              alt="nivas logo"
              class="logo"
              style="
                width: 107px;
                height: 30px;
                padding-right: 10px;

                transform: rotate(0deg);
                opacity: 1;
              "
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
</html>
`
