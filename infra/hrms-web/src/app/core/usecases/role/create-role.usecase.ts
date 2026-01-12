import { Inject, Injectable } from "@angular/core";
import { Role } from "../../domain/models/role.model";
import { ROLE_REPOSITORY, RoleRepository } from "../../domain/ports/role.repo";

@Injectable({ providedIn: 'root' })
export class CreateRoleUseCase {
    constructor(@Inject(ROLE_REPOSITORY) private roleRepo: RoleRepository) {}
    
    async Execute(role: Role): Promise<Role> {
        return this.roleRepo.create(role);
    }
}
