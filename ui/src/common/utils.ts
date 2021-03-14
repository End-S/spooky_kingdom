import { Subjects } from '@/common/models/article.model';

export function getSubjectEmoji(subject: Subjects): string {
  const subjectEmojis: { [ key: string]: string } = {
    [ Subjects.ALL ]: '',
    [ Subjects.GHOST ]: '👻 ',
    [ Subjects.UFO ]: '🛸',
    [ Subjects.WEIRD ]: '🐾',
  };
  return subjectEmojis[ subject ] || '?';
}
