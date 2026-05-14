const { execSync } = require('child_process');

/**
 * Executes BDD feature tests.
 * Usage: node run-bdd.cjs <path-to-feature-file>
 */
const featurePath = process.argv[2];
if (!featurePath) {
    console.error('Error: Please provide a path to a feature file.');
    process.exit(1);
}

try {
    // Assuming go-life uses a test runner that supports features, 
    // or we are running cucumber-js/godog. 
    // Adjust command as necessary for your environment.
    const cmd = `go test -v -tags=cucumber ./...`; 
    console.log(`Running tests for ${featurePath}...`);
    execSync(cmd, { stdio: 'inherit' });
} catch (error) {
    console.error('Test execution failed:', error.message);
    process.exit(1);
}
