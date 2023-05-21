package main

func (m Post) Comments() ([]*Comment, error) {
	comments, err := Comment{}.Where("post_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (m Post) JoinComments() ([]*Post, error) {
	var posts []*Post
	rows, err := db.Query(`
		select posts.id, posts.content, posts.author, posts.created_at, posts.updated_at from posts
			inner join comments
			on posts.id = comments.post_id;
	`)
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Content, &post.Author, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, err
}
