package models

/**
*
* request
* <p>
* request file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 13/12/2024
*
 */

type CreateEmailRequest struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Status string `json:"status" binding:"required"`
	//, oneof=USD
}
