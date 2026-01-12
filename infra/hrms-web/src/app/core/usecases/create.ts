import { Inject, Injectable } from "@angular/core";
import { CreateUser, UserData } from "../domain/models";
import { UserRepository, USER_REPOSITORY } from "../domain/ports/user.repo";

@Injectable({ providedIn: 'root' })
export class CreateUserUseCase {
    constructor(@Inject(USER_REPOSITORY) private userRepo: UserRepository) {}
    
    async Execute(user: CreateUser): Promise<UserData> {
        return this.userRepo.createUser(user);
    }
}
