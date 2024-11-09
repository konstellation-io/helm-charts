# GitHub Actions

* [`[GitHub] Auto-assign`](./github-auto-assign.yml): automatically assigns reviewers to pull requests when they are opened or marked as ready for review.

* [`[GitHub] Mark release as pre-release"`](./github-set-prerelease.yml): marks a release as a pre-release if its tag indicates it's a release candidate (example: `-rc.x`). This workflow ensures proper labeling of release candidates and integrates seamlessly with Helm Releaser actions.

* [`[GitHub] Stale issues and PRs`](./github-stale-issues-pr.yml): runs daily to identify and label issues or pull requests as stale if there has been no activity for 60 days. Stale items are optionally closed after a specified period, keeping the repository organized.

* [`[Helm] Check KDL Server major dependencies releases`](./helm-check-kdl-server-major-dependencies.yml): scheduled monthly or manually triggered, this workflow checks and updates the KDL Server Helm chart's major dependencies. It ensures the latest versions are applied and creates a pull request for the updates and improving dependency management.

* [`[Helm] Check KDL Server minor dependencies releases`](./helm-check-kdl-server-minor-dependencies.yml): runs weekly or manually to check and update minor dependencies for the KDL Server Helm chart. It adheres to semantic versioning rules and creates pull requests for updates, keeping the chart up-to-date.

* [`[Helm] Check KDL Server new releases`](./helm-check-kdl-server-release.yml): executed daily or manually to detect new KDL Server releases. If a new version is found, it updates the Helm chart and creates a pull request to reflect the change, including details about the release.

* [`[Helm Charts] Lint and test PR`](./helm-lint-test.yml): triggered by pull requests or manually to validate Helm charts. It performs linting, dependency checks and testing on Kubernetes clusters to ensure charts meet quality standards before merging.

* [`[Helm Charts] Releases`](./helm-release.yml): automatically generates and publishes Helm chart releases when changes are pushed to the main branch. The workflow signs and publishes OCI-compliant charts to the GitHub Container Registry, ensuring secure and traceable releases.
