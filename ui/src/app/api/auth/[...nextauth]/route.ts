import NextAuth from "next-auth";
import GoogleProvider from "next-auth/providers/google";

export const authOptions = {
    providers: [
        GoogleProvider({
            clientId: process.env.GOOGLE_CLIENT_ID || "",
            clientSecret: process.env.GOOGLE_CLIENT_SECRET || "",
        }),
    ],
    session: {
        strategy: "jwt" as const,
    },
    callbacks: {
        async jwt({ token, account }: any) {
            if (account) {
                token.accessToken = account.access_token;
            }
            return token;
        },
        async session({ session, token }: any) {
            // Send properties to the client
            session.accessToken = token.accessToken;
            // We also make the raw JWT token available so we can send it to the Go backend
            // In next-auth, the raw JWT is stored in cookies, but we can pass it explicitly or read it from the client.
            // NextAuth automatically signs the JWT, we'll see how to send it.
            return session;
        },
    },
};

const handler = NextAuth(authOptions);
export { handler as GET, handler as POST };
