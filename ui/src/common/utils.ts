import {
  capitalize,
  upperCase,
} from 'lodash-es';
import {
  Subjects,
  SubjectSelection,
} from '@/common/models/article.model';

export function getSubjectEmoji(subject: Subjects): string {
  const subjectEmojis: { [ key: string ]: string } = {
    [ Subjects.ALL ]: '',
    [ Subjects.GHOST ]: '👻 ',
    [ Subjects.UFO ]: '🛸',
    [ Subjects.ABC]: '🦁',
  };
  return subjectEmojis[ subject ] || '?';
}

export const availableSubjects: SubjectSelection[] = [
  {
    title: `${ getSubjectEmoji(Subjects.GHOST) } ${ capitalize(Subjects.GHOST) }`,
    value: Subjects.GHOST,
  },
  {
    title: `${ getSubjectEmoji(Subjects.UFO) }  ${ upperCase(Subjects.UFO) }`,
    value: Subjects.UFO,
  },
  {
    title: `${ getSubjectEmoji(Subjects.ABC) }  ${ upperCase(Subjects.ABC) }`,
    value: Subjects.ABC,
  },
];

export const filterSubjects: SubjectSelection[] = [
  { title: 'All', value: Subjects.ALL },
  ...availableSubjects,
];
