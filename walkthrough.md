# Walkthrough: Next.js + Go Backend Integration

The basic setup for a Next.js application with NextAuth social login (Google) and a Go backend is complete. You can now test it out!

## What was built

- A Next.js 15 application using Tailwind CSS and the App Router.
- NextAuth configured with the Google provider to issue JWTs.
- A custom Next.js landing page with a Google sign-in button.
- A dashboard page (accessible after login) that fetches data from the Go backend.
- A Go HTTP server using only the standard library (`net/http`) that:
  - Has CORS middleware allowing requests from `http://localhost:3000`.
  - Parses the `Authorization: Bearer <token>` header.
  - Decodes the JWT payload using standard `encoding/base64` to extract the user's email.
  - Returns a JSON response containing the email.

## How to Test

1. Add your Google OAuth credentials:
   Create a `.env.local` file in the `ui` directory with the following variables:
   ```env
   GOOGLE_CLIENT_ID=your-google-client-id
   GOOGLE_CLIENT_SECRET=your-google-client-secret
   NEXTAUTH_SECRET=a-random-secret-string
   NEXTAUTH_URL=http://localhost:3000
   ```

2. Start the Go server in one terminal:
   ```bash
   cd /home/rick/Projects/custom-start-bible-reading-plan
   go run cmd/server/main.go
   ```

3. Start the Next.js dev server in another terminal:
   ```bash
   cd /home/rick/Projects/custom-start-bible-reading-plan/ui
   npm run dev
   ```

4. Navigate to `http://localhost:3000` in your browser.
5. Click "Sign in with Google" and authenticate.
6. Once redirected back, you will see your profile information.
7. Click the "Call Go Backend" button to make an authenticated request.
8. The Go backend will decode your JWT, extract your email, and return it in a JSON response seamlessly!
