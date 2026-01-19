import { Inject, Injectable } from "@angular/core";
import { Position, UpdatePositionDto } from "../../domain/models/position.model";
import { POSITION_REPOSITORY, PositionRepository } from "../../domain/ports/position.repo";

@Injectable({ providedIn: 'root' })
export class UpdatePositionUseCase {
    constructor(@Inject(POSITION_REPOSITORY) private positionRepo: PositionRepository) {}
    
    async Execute(position: UpdatePositionDto): Promise<Position> {
        return this.positionRepo.update(position);
    }
}
