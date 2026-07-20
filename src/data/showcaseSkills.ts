export type ShowcaseSkill = {
  slug: string;
  name: string;
  title: string;
  description: string;
  content: string;
};

const skillFiles = import.meta.glob(
  '../../showcase-skills/*/skill/*/SKILL.md',
  {
    eager: true,
    import: 'default',
    query: '?raw',
  },
) as Record<string, string>;

function readFrontmatterValue(content: string, key: string) {
  const frontmatter = content.match(/^---\n([\s\S]*?)\n---/u)?.[1];
  const line = frontmatter
    ?.split('\n')
    .find((candidate) => candidate.startsWith(`${key}:`));

  return line?.slice(key.length + 1).trim() ?? '';
}

function formatTitle(name: string) {
  return name
    .split('-')
    .map((word) => `${word.charAt(0).toUpperCase()}${word.slice(1)}`)
    .join(' ');
}

export const showcaseSkills = Object.entries(skillFiles)
  .map(([path, content]) => {
    const slug = path.match(/showcase-skills\/([^/]+)\//u)?.[1] ?? path;
    const name = readFrontmatterValue(content, 'name') || slug;

    return {
      slug,
      name,
      title: formatTitle(name),
      description: readFrontmatterValue(content, 'description'),
      content,
    } satisfies ShowcaseSkill;
  })
  .sort((left, right) => left.title.localeCompare(right.title));
