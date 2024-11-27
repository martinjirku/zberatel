import { useQuery } from "@apollo/client";
import { graphql } from "./gql";
import DefaultLayout from "./layouts/default-layout";

export const MY_COLLECTIONS = graphql(`
  query MyCollections($input: CollectionsListInput!) {
    myCollectionsList(input: $input) {
      items {
        id
        title
        description
      }
      meta {
        total
        nextPage {
          limit
          offset
        }
        currentPage {
          limit
          offset
        }
      }
    }
  }
`);

function App() {
  const {} = useQuery(MY_COLLECTIONS, {
    variables: { input: { paging: { limit: 4, offset: 0 } } },
  });
  return (
    <DefaultLayout>
      <main className="flex-1 bg-gray-50 overflow-y-auto p-4">
        <h1 className="text-2xl font-bold mb-4">Welcome to Kulektor</h1>
        <p className="text-gray-700">
          This is your main content area. Add components, cards, or anything you
          like here.
        </p>
      </main>
    </DefaultLayout>
  );
}

export default App;
