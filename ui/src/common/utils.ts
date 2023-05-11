import { Subjects } from "@/common/models/article.model";
import { capitalize, upperCase } from "lodash-es";
import { DateTime, type DateTimeFormatOptions } from "luxon";
import type { SubjectSelection } from "./models/article.model";

export function getSubjectEmoji(subject: Subjects): string {
  const subjectEmojis: { [key: string]: string } = {
    [Subjects.ALL]: "",
    [Subjects.GHOST]: "👻 ",
    [Subjects.UFO]: "🛸",
    [Subjects.CRYPTID]: "🦄",
  };
  return subjectEmojis[subject] || "?";
}

export const availableSubjects: SubjectSelection[] = [
  {
    title: `${getSubjectEmoji(Subjects.GHOST)} ${capitalize(Subjects.GHOST)}`,
    value: Subjects.GHOST,
  },
  {
    title: `${getSubjectEmoji(Subjects.UFO)}  ${upperCase(Subjects.UFO)}`,
    value: Subjects.UFO,
  },
  {
    title: `${getSubjectEmoji(Subjects.CRYPTID)}  ${capitalize(Subjects.CRYPTID)}`,
    value: Subjects.CRYPTID,
  },
];

export const filterSubjects: SubjectSelection[] = [
  { title: "All", value: Subjects.ALL },
  ...availableSubjects,
];

export function formatIso(IsoTime: string, format: DateTimeFormatOptions): string {
  return DateTime.fromISO(IsoTime).toLocaleString(format);
}
