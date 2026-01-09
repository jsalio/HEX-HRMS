export enum UserType {
  Admin = 'admin',
  Normal = 'normal'
}

export interface User {
  id: string;
  username: string;
  password: string;
  email: string;
  type: UserType;
  active: boolean;
}

export interface CreateUser {
  username: string;
  password: string;
  email: string;
  type: UserType;
}

export interface ModifyUser {
  id: string;
  username: string;
  password: string;
  email: string;
  type: UserType;
}

export interface UserData {
  id: string;
  username: string;
  email: string;
  type: UserType;
}

export interface LoginUser {
  username: string;
  password: string;
}
