import { Inject, Injectable } from "@angular/core";
import { Filter, LoginUser, SearchQuery, UserData } from "../domain/models";
import { UserRepository, USER_REPOSITORY } from "../domain/ports/user.repo";

@Injectable({ providedIn: 'root' })

export class ListUserUseCase {
    /**
     *
     */
    constructor(@Inject(USER_REPOSITORY) private userRepo: UserRepository) {}
    
    async Execute(query: SearchQuery): Promise<UserData[]> {
        const data = await this.userRepo.list(query);
        return data
    }
}