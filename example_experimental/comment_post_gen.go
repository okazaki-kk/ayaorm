package main

func (m Post) Comments() ([]*Comment, error) {
	comments, err := Comment{}.Where("post_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
