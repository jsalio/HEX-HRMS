import { Inject, Injectable } from "@angular/core";
import { POSITION_REPOSITORY, PositionRepository } from "../../domain/ports/position.repo";

@Injectable({ providedIn: 'root' })
export class DeletePositionUseCase {
    constructor(@Inject(POSITION_REPOSITORY) private positionRepo: PositionRepository) {}
    
    async Execute(id: string): Promise<void> {
        return this.positionRepo.delete(id);
    }
}
