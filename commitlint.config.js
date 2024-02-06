module.exports = {
    extends: ['@commitlint/config-conventional'],
    rules: {
      'type-enum': [1, 'always', ['ci', 'chore', 'docs', 'feat', 'fix', 'perf', 'refactor', 'revert', 'test', 'wip']],
      'subject-case': [1, 'always', 'sentence-case'],
    },
  };
