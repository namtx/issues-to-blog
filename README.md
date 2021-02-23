# issues-to-blog

This action turn Github repository's issues into Jekyll blog post and post into another repository (same user)

## Inputs

### `targetRepository`

**Required** The repository where to put the blog markdown file

### `targetBranch`

**Required** The branch where to put the blog markdown file

### `targetDirectory`

**Required** The directory where to put the blog markdown file

### `sourceRepository`

**Required** The repository where contains issues

### `accessToken`

**Required** Access Token is required to push new commit cross Github repositories

## Example usage

```yml
steps:
  - name: Issues to blog post action
    uses: namtx/issues-to-blog@0.0.2
      with:
	targetRepository: thinkspace
	targetBranch: master
	targetDirectory: _posts
	sourceRepository: til
	accessToken: ${{ secrets.ACCESS_TOKEN }}
```
