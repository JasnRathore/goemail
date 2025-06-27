
# GoEmail 📧

A lightweight, developer-friendly Go package for sending HTML emails with attachment support and built-in tracking capabilities. Perfect for applications that need reliable email delivery without the complexity of larger email libraries.

## Project Requirements

- **Go Version**: 1.24.4 or higher
- **SMTP Server Access**: Valid SMTP credentials for your email provider
- **Network Access**: Ability to connect to SMTP servers on standard ports (25, 587, etc.)

## Dependencies

This package uses minimal external dependencies to keep things lightweight:

```go
gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc
```

All dependencies are automatically managed through Go modules.

## Getting Started

### Installation

Add the package to your project:

```bash
go get github.com/jasnrathore/goemail
```

### Basic Setup

Import the package and create your first mail profile:

```go
import "github.com/jasnrathore/goemail"

profile := goemail.NewProfile(
    "Marketing Bot",                    // Display name
    "your-email@gmail.com",            // SMTP username
    "Marketing Bot <your-email@gmail.com>", // From header
    "smtp.gmail.com:587",              // SMTP host:port
    "your-app-password",               // SMTP password
)
```

## How to Run the Application

### Sending Basic Emails

The simplest way to send an email:

```go
err := profile.SendMail(
    "recipient@example.com",
    "Welcome to Our Service",
    "<h1>Hello!</h1><p>Welcome to our platform.</p>",
    nil, // no attachments
)

if err != nil {
    log.Printf("Failed to send email: %v", err)
}
```

### Adding Attachments

Send emails with file attachments:

```go
attachments := []goemail.MailAttachment{
    {
        FileName: "report.pdf",
        Data:     pdfData, // []byte containing file data
    },
    {
        FileName: "logo.png",
        Data:     imageData,
    },
}

err := profile.SendMail(
    "client@company.com",
    "Monthly Report",
    "<p>Please find attached your monthly report.</p>",
    attachments,
)
```

### Email Tracking

Add invisible tracking pixels to monitor email opens:

```go
trackingURL := "https://your-analytics-server.com/track?user=123&campaign=welcome"

err := profile.SendMailWithTracking(
    "user@example.com",
    "Welcome Email",
    "<h2>Welcome aboard!</h2><p>We're excited to have you.</p>",
    nil,
    trackingURL,
)
```

The tracking image is automatically appended as a 1x1 pixel invisible element.

### Testing Your Configuration

Before sending real emails, test your SMTP settings:

```go
err := profile.SendTestMail("test-recipient@example.com")
if err != nil {
    log.Printf("SMTP configuration issue: %v", err)
} else {
    log.Println("Test email sent successfully!")
}
```

## Relevant Examples

### Complete Email Campaign Example

```go
package main

import (
    "log"
    "github.com/jasnrathore/goemail"
)

func main() {
    // Configure your email profile
    profile := goemail.NewProfile(
        "Newsletter Bot",
        "newsletter@yourcompany.com",
        "Company Newsletter <newsletter@yourcompany.com>",
        "smtp.gmail.com:587",
        "your-secure-app-password",
    )

    recipients := []string{
        "subscriber1@example.com",
        "subscriber2@example.com",
    }

    htmlBody := `
        <html>
        <body>
            <h2>This Week in Tech</h2>
            <p>Here are the latest updates from our team...</p>
            <img src="https://yoursite.com/newsletter-banner.jpg" alt="Banner" />
        </body>
        </html>
    `

    for _, recipient := range recipients {
        trackingURL := fmt.Sprintf("https://analytics.yoursite.com/open?email=%s", recipient)
        
        err := profile.SendMailWithTracking(
            recipient,
            "This Week in Tech - Newsletter #42",
            htmlBody,
            nil,
            trackingURL,
        )

        if err != nil {
            log.Printf("Failed to send to %s: %v", recipient, err)
        } else {
            log.Printf("Newsletter sent to %s", recipient)
        }
    }
}
```

### Error Handling Best Practices

```go
func sendWelcomeEmail(profile goemail.MailProfile, userEmail string) error {
    err := profile.SendMail(
        userEmail,
        "Welcome to Our Platform",
        generateWelcomeHTML(),
        nil,
    )

    if err != nil {
        // Log the error for debugging
        log.Printf("Email delivery failed for %s: %v", userEmail, err)
        
        // You might want to queue for retry or alert administrators
        return fmt.Errorf("welcome email delivery failed: %w", err)
    }

    return nil
}
```

### Working with Different SMTP Providers

```go
// Gmail configuration (TLS/STARTTLS)
gmailProfile := goemail.NewProfile(
    "App Name",
    "yourapp@gmail.com",
    "App Name <yourapp@gmail.com>",
    "smtp.gmail.com:587",
    "your-app-password", // Use App Password, not regular password
)

// Outlook/Hotmail configuration (TLS/STARTTLS)
outlookProfile := goemail.NewProfile(
    "App Name",
    "yourapp@outlook.com",
    "App Name <yourapp@outlook.com>",
    "smtp-mail.outlook.com:587",
    "your-password",
)

// Yahoo Mail configuration (SSL)
yahooProfile := goemail.NewProfile(
    "App Name",
    "yourapp@yahoo.com",
    "App Name <yourapp@yahoo.com>",
    "smtp.mail.yahoo.com:465",
    "your-app-password",
)

// SendGrid configuration
sendgridProfile := goemail.NewProfile(
    "App Name",
    "apikey", // Username is literally "apikey"
    "App Name <verified-sender@yourdomain.com>",
    "smtp.sendgrid.net:587",
    "your-sendgrid-api-key",
)

// Custom SMTP server (port varies by provider)
customProfile := goemail.NewProfile(
    "Corporate Mail",
    "noreply@yourcompany.com",
    "Company Name <noreply@yourcompany.com>",
    "mail.yourcompany.com:25", // Could be 25, 587, 465, or 2525
    "smtp-password",
)
```

## Running Tests

Execute the test suite to verify functionality:

```bash
go test -v
```

**Note**: Some tests require valid SMTP credentials. Update the test configuration in `goemail_test.go` with your actual email settings for comprehensive testing.

## Configuration Tips

### Gmail Setup
1. Generate an App Password specifically for this application
2. Use the App Password instead of your regular Gmail password

### Security Considerations
- Store SMTP passwords in environment variables, not in source code
- Use dedicated email accounts for automated sending
