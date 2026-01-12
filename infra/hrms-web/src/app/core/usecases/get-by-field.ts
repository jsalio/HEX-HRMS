import { Inject, Injectable } from "@angular/core";
import { Filter, UserData } from "../domain/models";
import { USER_REPOSITORY, UserRepository } from "../domain/ports/user.repo";

@Injectable({
    providedIn: 'root'
})
export class GetUserByFieldUseCase {
    constructor(@Inject(USER_REPOSITORY) private userRepository: UserRepository) {}

    Execute(filter: Filter): Promise<UserData> {
        return this.userRepository.getUserByFilter(filter);
    }
}
