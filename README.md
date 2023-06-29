# Github-Weekly-Repo
A GitHub integration basically, it sends weekly emails on the top repositories 

# Requirements
Yes, there are a few placeholders in the code that you need to replace with your own information:

1. In `main.go`:
   - Replace `"YOUR_GITHUB_ACCESS_TOKEN"` with your GitHub personal access token.
   - Replace `"your-email@example.com"` with the recipient email address where you want to receive the weekly top GitHub repositories.
   - Replace `"sender-email@example.com"` with the sender email address.
   - Replace `"smtp.example.com"` with your SMTP server address.
   - Replace `587` with the port number of your SMTP server.
   - Replace `"smtp-user"` with your SMTP server username.
   - Replace `"smtp-password"` with your SMTP server password.

2. In `email_sender.go`:
   - Replace `"sender-email@example.com"` with the sender email address. This should be the same as the one used in `github_integration.go`.

Ensure that you have a valid GitHub personal access token and access to an SMTP server for sending emails.

Once you have replaced the placeholders with your information, save the files and follow the next steps to run the integration.

3. Open a terminal or command prompt and navigate to the directory containing the `github_integration.go` file.

4. Run the following command to execute the integration:
   ```
   go run github_integration.go
   ```

   This will fetch the weekly top GitHub repositories and send an email with the results.

5. By default, the email sending functionality is commented out in the `sendEmail` function. Uncomment the line `err := d.DialAndSend(m)` to enable sending emails.

   **Note:** Be cautious when uncommenting this line to avoid accidentally sending emails during testing. Uncomment it only when you are ready to send emails.

That's it! You have set up and run the integration system. Make sure to review the code and customize it further according to your specific requirements and preferences.
