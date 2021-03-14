export interface ArticleGetParams {
  pbs?: string[];
  frm?: number;
  to?: number;
  ord: Ordering;
  srt: ArticleSortBy;
  sbj?: Subjects;
  oft: number;
  lmt: 10 | 25 | 50;
  pnd?: boolean;
}

export interface Pagination {
  currentPage: number;
  pageSize: 10 | 25 | 50;
}

export interface ArticleFilters {
  sortBy: ArticleSortBy;
  order: Ordering;
  subject: Subjects;
  publishers: string[];
  date: {
    from?: Date;
    to?: Date;
  };
}

export interface Article {
  id: string;
  publisher: {
    id: string;
    name: string;
  };
  datePublished: string;
  dateRetrieved: string;
  title: string;
  description: string;
  link: string;
  type: Subjects;
  accepted: boolean;
}

export enum Subjects {
  ALL = '',
  GHOST = 'ghost',
  UFO = 'ufo',
  WEIRD = 'weird',
}

export enum ArticleSortBy {
  DATE = 'date',
  TITLE = 'title',
  DESCRIPTION = 'description',
}

export enum Ordering {
  ASCENDING = 'asc',
  DESCENDING = 'desc',
}
