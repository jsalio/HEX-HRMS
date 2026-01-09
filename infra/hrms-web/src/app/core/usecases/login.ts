import { Inject } from "@angular/core";
import { USER_REPOSITORY } from "../../app.config";
import { LoginUser, UserData } from "../domain/models";
import { UserRepository } from "../domain/ports/user.repo";

export class LoginUserUseCase {
    /**
     *
     */
constructor(@Inject(USER_REPOSITORY) private userRepo: UserRepository) {}
    async Execute(loginUser:LoginUser): Promise<UserData> {
        const data = await this.userRepo.login(loginUser);
        return data
    }
}0