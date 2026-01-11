export interface PaginatedResponse<T> {
    total_rows: number;
    total_pages: number;
    rows: T[];
}
