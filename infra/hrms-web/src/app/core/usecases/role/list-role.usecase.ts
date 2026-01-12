import { Inject, Injectable } from "@angular/core";
import { PaginatedResponse, SearchQuery } from "../../domain/models";
import { Role } from "../../domain/models/role.model";
import { ROLE_REPOSITORY, RoleRepository } from "../../domain/ports/role.repo";

@Injectable({ providedIn: 'root' })
export class ListRoleUseCase {
    constructor(@Inject(ROLE_REPOSITORY) private roleRepo: RoleRepository) {}
    
    async Execute(query: SearchQuery): Promise<PaginatedResponse<Role>> {
        return this.roleRepo.list(query);
    }
}
