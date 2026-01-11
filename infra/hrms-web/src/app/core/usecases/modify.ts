import { Inject, Injectable } from "@angular/core";
import { ModifyUser } from "../domain/models";
import { USER_REPOSITORY, UserRepository } from "../domain/ports/user.repo";

@Injectable({
    providedIn: 'root'
})
export class ModifyUserUseCase {
    constructor(@Inject(USER_REPOSITORY) private userRepository: UserRepository) {}

    Execute(user: ModifyUser): Promise<any> {
        return this.userRepository.updateUser(user);
    }
}
