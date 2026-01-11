export interface Filter {
  key: string;
  value: any;
}

export type Filters = Filter[];

export interface Pagination {
  page: number;
  limit: number;
}

export interface SearchQuery {
  filters: Filters;
  pagination: Pagination;
}
