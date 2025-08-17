export interface User {
  id: string;
  email: string;
  username: string;
  metadata: {
    name?: string;
    given_name?: string;
    family_name?: string;
    picture?: string;
    updated_at?: string;
  };
}

export interface AuthStatus {
  authenticated: boolean;
  user?: User;
  message?: string;
}
