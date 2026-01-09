import { CreateUser, Filter, LoginUser, ModifyUser, User, UserData } from "../models";

export interface UserRepository {
    // getUserById(id: string): Promise<User | null>;
    // getUserByFilter(filter: Filter): Promise<User | null>;
    // getAllUsers(): Promise<User[]>;
    // createUser(user: CreateUser): Promise<User>;
    // updateUser(id: string, user: ModifyUser): Promise<User>;
    // deleteUser(id: string): Promise<void>;
    login(loginUser:LoginUser):Promise<UserData>
}