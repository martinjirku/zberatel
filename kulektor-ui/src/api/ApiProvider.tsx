import { ReactElement, ReactNode, useMemo } from "react";
import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  HttpLink,
  from,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { onError } from "@apollo/client/link/error";
import { useAuth0 } from "@auth0/auth0-react";

export default function ApiProvider({
  children,
}: {
  children: ReactNode;
}): ReactElement {
  const { getAccessTokenSilently } = useAuth0();
  const client = useMemo(() => {
    const authLink = setContext(async (_, { headers }) => {
      const token = await getAccessTokenSilently(); // Call getAccessTokenSilently function
      return {
        headers: {
          ...headers,
          authorization: token ? `Bearer ${token}` : "", // Use the bearer schema
        },
      };
    });
    const httpLink = new HttpLink({
      uri: import.meta.env.VITE_GRAPHQL_API || "http://localhost:3000/graphql",
    });
    const errorLink = onError(({ graphQLErrors, networkError }) => {
      console.log(">>> ERROR API:", graphQLErrors, networkError);
    });

    return new ApolloClient({
      link: from([authLink, errorLink, httpLink]),
      cache: new InMemoryCache(),
    });
  }, []);
  return <ApolloProvider client={client}>{children}</ApolloProvider>;
}
