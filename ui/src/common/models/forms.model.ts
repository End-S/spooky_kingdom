import type { ArticleState } from "@/stores/article.store";
import type { ArticleFilters } from "./article.model";

export type FilterFormValue = ArticleFilters & Pick<ArticleState, "pageSize" | "dateSpan">;
