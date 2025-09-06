export const LOGIN_PATH = "btmwfofxsa"

export interface UserAuthForm {
  username: string;
  password: string;
}

export interface GenericToken<D> {
  id?: string;
  userId?: number;
  accessToken?: string;
  refreshToken?: string;
  expiresAt?: D;
  createdAt?: D;
  updatedAt?: D;
}

export interface TokenApi extends GenericToken<string> {
}

export interface Token extends GenericToken<Date> {
}

export interface User {
  id?: number;
  details?: UserDetails;
  username?: string;
  verified?: boolean;
}

export interface UserDetails {
  firstName?: string;
  lastName?: string;
  email?: string;
  dateOfBirth?: string;
}
