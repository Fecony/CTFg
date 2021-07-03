import { gql } from "@apollo/client";
import { AuthContextQuery, useAuthContextQuery } from "generated";
import React from "react";

gql`
  query AuthContext {
    current_user {
      id
      name
      created
      country
      affiliation
      bracket
      website
      verified
      type
      captain_of {
        id
      }
      team {
        id
        name
        created
        country
        affiliation
        bracket
        website
      }
    }
    config(where: { key: { _in: ["ctf_name"] } }) {
      key
      value
    }
  }
`;

type AuthContextType = {
  loading: boolean;
  current_user?: AuthContextQuery["current_user"][0];
  config: { [key: string]: string };
};

export const AuthContext = React.createContext<AuthContextType>({
  loading: true,
  config: {},
});

type ChildrenProps = {
  children?: React.ReactNode;
};

export function useAuthContext() {
  return React.useContext(AuthContext);
}

export function AuthContextProvider({ children }: ChildrenProps) {
  const { data, loading, error } = useAuthContextQuery();

  if (loading || error) {
    console.log(error);
    return <>{children}</>;
  }

  return (
    <AuthContext.Provider
      value={{
        loading,
        current_user: data?.current_user[0],
        // @ts-ignore this doesnt work
        config: Object.fromEntries(data.config),
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
