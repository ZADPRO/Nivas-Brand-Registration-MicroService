package BrandInviteMailTemplate

var InviteTemplate = `<!DOCTYPE html>
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
        margin: 20px 20px 15px 15px;
        background: #ffffff;
        border-radius: 8px;
        overflow: hidden;
        border: 1;
        border: #f0eeec;
        justify-content: center;
        text-align: justify;

      }
      .header {
        text-align: start;
        padding: 20px;
      }
      .header img.logo {
        width: 105px;
        height: 32;
        opacity: 1;
        margin: 10px 10px 20px 0px;
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
        <h2>Start Your Journey with us</h2>
        <p>
          Step into an exclusive marketplace designed to elevate your brand and
          connect you with discerning customers.
        </p>
      </div>

      <!-- Hero Image -->
      <div class="hero">
        <img src="https://raw.githubusercontent.com/vijayzadroit/nivas/main/inviteHero.png" alt="Banner" />
      </div>

      <!-- Content -->
      <div class="container2">
        <div class="content">
          <span style="font-size: 18px; font-weight: 700">
            Dear {{.ContactPerson}},
          </span>
          <p>I hope this message finds you well.</p>
          <p>
            We at <strong>Nivas - House of Celebrities</strong> are excited to
            extend an exclusive invitation for <strong>{{.BrandName}}</strong> to
            join our curated marketplace that celebrates and promotes
            celebrity-owned brands. Our platform is designed to offer customers
            a premium shopping experience by providing authentic products
            created by renowned personalities.
          </p>

          <p>
            By partnering with Nivas - House of Celebrities, your brand will
            gain access to:
          </p>
          <ul>
            <li>
              A targeted audience of engaged consumers who value authenticity
              and celebrity associations.
            </li>
            <li>Dedicated brand storefront with high visibility.</li>
            <li>
              Seamless integration and dedicated support throughout the
              onboarding and selling process.
            </li>
            <li>
              Marketing initiatives highlighting your brand’s unique story and
              connection with the celebrity.
            </li>
          </ul>

          <p>
            We believe <strong>{{.BrandName}}</strong> perfectly aligns with our
            vision of delivering premium, aspirational products to consumers,
            and we are honored to have you on board.
          </p>

          <p>
            To get started please click on this link <a href="{{.RegistrationURL}}">Register Link</a> and register
            the brand. Once you have submitted the registration form, our team
            will review and get back to you within 3-4 working days.
          </p>

          <a href="{{.RegistrationURL}}" class="btn">Get started</a>
        </div>

        <!-- Footer -->
        <div class="footer">
          <p>Looking forward to the opportunity of collaborating with you.</p>
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
</html>`
