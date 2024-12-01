import {
  Auth0ContextInterface,
  useAuth0,
  User as Auth0User,
} from "@auth0/auth0-react";
import {
  createContext,
  FC,
  PropsWithChildren,
  useContext,
  useEffect,
} from "react";
// import jwtDecode from "jwt-decode";
import { Role } from "../gql/graphql";

export const RolesKey = "kulektor/roles";

export interface AuthState {
  isAuthenticated: boolean;
  loginWithPopup: Auth0ContextInterface<Auth0User>["loginWithPopup"];
  logout: Auth0ContextInterface<Auth0User>["logout"];
  user?: User;
  isCollector: () => boolean;
  isAdmin: () => boolean;
}

export interface User {
  userId: string;
  name?: string;
  username?: string;
  email?: string;
  firstName?: string;
  lastName?: string;
  picture?: string;
  roles: Role[];
}

const authContext = createContext<AuthState>({
  isAuthenticated: false,
  loginWithPopup: () => Promise.resolve(),
  logout: () => Promise.resolve(),
  isCollector: () => false,
  isAdmin: () => false,
});

export const AuthStateProvider: FC<PropsWithChildren> = ({ children }) => {
  const {
    loginWithPopup,
    isAuthenticated,
    logout,
    user: auth0User,
    getAccessTokenSilently,
  } = useAuth0<Auth0User>();

  useEffect(() => {
    getAccessTokenSilently();
  });
  let user: User | undefined;
  if (auth0User) {
    user = {
      userId: auth0User.userId,
      name: auth0User.name,
      username: auth0User.preferred_username,
      firstName: auth0User.given_name,
      lastName: auth0User.family_name,
      email: auth0User.email,
      picture: auth0User.picture,
      roles: auth0User[RolesKey] ?? [],
    };
    user.roles.push();
  }
  const state: AuthState = {
    isAuthenticated,
    loginWithPopup,
    logout,
    user,
    isCollector: () => !!user?.roles.includes(Role.Collector),
    isAdmin: () => !!user?.roles.includes(Role.Admin),
  };
  return <authContext.Provider value={state}>{children}</authContext.Provider>;
};

export const useAuth = () => {
  return useContext(authContext);
};