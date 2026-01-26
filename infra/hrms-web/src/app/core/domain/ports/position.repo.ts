import { InjectionToken } from "@angular/core";
import { PaginatedResponse, SearchQuery } from "../models";
import { Position, CreatePositionDto, UpdatePositionDto } from "../models/position.model";

/**
 * Position repository interface
 */
export interface PositionRepository {
  /**
   * List positions with pagination and filters
   */
  list(query: SearchQuery): Promise<PaginatedResponse<Position>>;

  /**
   * Get a position by ID
   */
  get(id: string): Promise<Position>;

  /**
   * Create a new position
   */
  create(position: CreatePositionDto): Promise<Position>;

  /**
   * Update an existing position
   */
  update(position: UpdatePositionDto): Promise<Position>;

  /**
   * Delete a position by ID
   */
  delete(id: string): Promise<void>;
}

/**
 * Injection token for Position repository
 */
export const POSITION_REPOSITORY = new InjectionToken<PositionRepository>('position.repository');
