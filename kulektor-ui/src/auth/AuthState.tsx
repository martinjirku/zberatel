import { Auth0ContextInterface, useAuth0, User } from "@auth0/auth0-react";
import {
  createContext,
  FC,
  PropsWithChildren,
  useContext,
  useEffect,
} from "react";

export interface AuthState {
  isAuthenticated: boolean;
  loginWithPopup: Auth0ContextInterface<User>["loginWithPopup"];
  logout: Auth0ContextInterface<User>["logout"];
  user?: User;
}

const authContext = createContext<AuthState>({
  isAuthenticated: false,
  loginWithPopup: () => Promise.resolve(),
  logout: () => Promise.resolve(),
});

export const AuthStateProvider: FC<PropsWithChildren> = ({ children }) => {
  const {
    loginWithPopup,
    isAuthenticated,
    logout,
    user,
    getAccessTokenSilently,
  } = useAuth0<User>();

  useEffect(() => {
    getAccessTokenSilently();
  });

  const state: AuthState = {
    isAuthenticated,
    loginWithPopup,
    logout,
    user,
  };
  return <authContext.Provider value={state}>{children}</authContext.Provider>;
};

export const useAuth = () => {
  return useContext(authContext);
};
