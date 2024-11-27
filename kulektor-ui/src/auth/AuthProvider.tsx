import React from "react";
import { Auth0Provider } from "@auth0/auth0-react";
import { AuthStateProvider } from "./AuthState";

export interface Auth0ProviderWithHistoryProps {
  children: React.ReactNode;
}

const AuthProvider = ({
  children,
}: Auth0ProviderWithHistoryProps): React.ReactElement => {
  // Retrieve the previously created environment variables
  const domain = import.meta.env.VITE_REACT_APP_AUTH0_DOMAIN;
  const clientId = import.meta.env.VITE_REACT_APP_AUTH0_CLIENT_ID;
  const audience = import.meta.env.VITE_REACT_APP_AUTH0_AUDIENCE;

  // Fail fast if the environment variables aren't set
  if (!domain || !clientId)
    throw new Error(
      "Please set VITE_REACT_APP_AUTH0_DOMAIN and VITE_REACT_APP_AUTH0_CLIENT_ID env. variables",
    );

  return (
    <Auth0Provider
      domain={domain}
      clientId={clientId}
      authorizationParams={{
        audience: audience,
        redirect_uri: window.location.origin,
      }}
      useRefreshTokens
      cacheLocation="localstorage"
    >
      <AuthStateProvider>{children}</AuthStateProvider>
    </Auth0Provider>
  );
};

export default AuthProvider;
