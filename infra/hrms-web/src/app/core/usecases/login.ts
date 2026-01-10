import { Inject, Injectable } from "@angular/core";
import { LoginUser, UserData } from "../domain/models";
import { UserRepository, USER_REPOSITORY } from "../domain/ports/user.repo";

@Injectable({ providedIn: 'root' })

export class LoginUserUseCase {
    /**
     *
     */
    constructor(@Inject(USER_REPOSITORY) private userRepo: UserRepository) {}
    
    async Execute(loginUser:LoginUser): Promise<UserData> {
        const data = await this.userRepo.login(loginUser);
        return data
    }
}