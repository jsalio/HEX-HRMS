import { Inject, Injectable } from "@angular/core";
import { Filter, LoginUser, UserData } from "../domain/models";
import { UserRepository, USER_REPOSITORY } from "../domain/ports/user.repo";

@Injectable({ providedIn: 'root' })

export class ListUserUseCase {
    /**
     *
     */
    constructor(@Inject(USER_REPOSITORY) private userRepo: UserRepository) {}
    
    async Execute(filter: Filter): Promise<UserData[]> {
        const data = await this.userRepo.list(filter);
        return data
    }
}