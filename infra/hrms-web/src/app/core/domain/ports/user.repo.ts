import { InjectionToken } from "@angular/core";
import { CreateUser, Filter, LoginUser, ModifyUser, SearchQuery, User, UserData } from "../models";

export const USER_REPOSITORY = new InjectionToken<UserRepository>('USER_REPOSITORY');

export interface UserRepository {
    // getUserById(id: string): Promise<User | null>;
    // getUserByFilter(filter: Filter): Promise<User | null>;
    // getAllUsers(): Promise<User[]>;
    // createUser(user: CreateUser): Promise<User>;
    // updateUser(id: string, user: ModifyUser): Promise<User>;
    // deleteUser(id: string): Promise<void>;
    login(loginUser:LoginUser):Promise<UserData>
    list(query: SearchQuery):Promise<UserData[]>
}