export enum UserType {
  Admin = 'admin',
  Normal = 'normal'
}

export interface User {
  id: string;
  username: string;
  name: string;
  lastName: string;
  password: string;
  email: string;
  type: UserType;
  active: boolean;
}

export interface CreateUser {
  username: string;
  name: string;
  lastName: string;
  password: string;
  email: string;
  type: UserType;
}

export interface ModifyUser {
  id: string;
  username: string;
  name: string;
  lastName: string;
  password: string;
  email: string;
  type: UserType;
}

export interface UserData {
  id: string;
  username: string;
  name: string;
  lastName: string;
  email: string;
  type: UserType;
  token: string;
  picture?: string;
  role?: string;
}

export interface LoginUser {
  username: string;
  password: string;
}
