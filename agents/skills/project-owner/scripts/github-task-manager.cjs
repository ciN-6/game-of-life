const { execSync } = require('child_process');
const fs = require('fs');

/**
 * Creates a GitHub issue from a task template file.
 * Usage: node github-task-manager.cjs <path-to-task-template>
 */
const filePath = process.argv[2];
if (!filePath || !fs.existsSync(filePath)) {
    console.error('Error: Provide a valid path to a task template file.');
    process.exit(1);
}

const content = fs.readFileSync(filePath, 'utf8');

// Parse Title
const titleMatch = content.match(/^Title: (.*)/m);
const title = titleMatch ? titleMatch[1] : "New Task";

// Parse Labels
const labelsMatch = content.match(/^Labels: (.*)/m);
const labels = labelsMatch ? labelsMatch[1].split(',').map(l => l.trim()).join(',') : "";

// Use the full file content as the body
const body = content;

try {
    // Escape double quotes for shell command
    const escapedBody = body.replace(/"/g, '\\"');
    const cmd = `gh issue create --title "${title}" --body "${escapedBody}" ${labels ? `--label "${labels}"` : ""}`;
    execSync(cmd, { stdio: 'inherit' });
    console.log(`Successfully created issue: ${title}`);
} catch (error) {
    console.error('Failed to create issue:', error.message);
}
