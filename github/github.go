package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

const GitLabAPI = "https://gitlab.com/api/v4/projects"

type Client struct {
	Token string
	hc    http.Client
}
type ProjectBasicDetails struct {
	ProjectId              int    `json:"id"`
	ProjectName            string `json:"name"`
	ProjectDescription     string `json:"description"`
	ProjectCreatedAt       string `json:"created_at"`
	ProjectLastActivityAt  string `json:"last_activity_at"`
	ProjectWebUrl          string `json:"web_url"`
	ProjectAvatarUrl       string `json:"avatar_url"`
	ProjectStarCount       int    `json:"star_count"`
	ProjectForksCount      int    `json:"forks_count"`
	ProjectOpenIssuesCount int    `json:"open_issues_count"`
}
type ProjectLanguage struct {
	Language map[string]float64 `json:"languages"`
}
type Members struct {
	UserId    int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	WebUrl    string `json:"web_url"`
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

func (c *Client) requestDoWithAuth(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	resp, err := c.hc.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func GetProjectDetails(c *fiber.Ctx) error {
	project_id, err := c.ParamsInt("id")
	var token = os.Getenv("GITLAB_TOKEN")

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Please pass valid project id to get the data",
		})
	}

	var client = NewClient(token)

	url := fmt.Sprintf(GitLabAPI+"/%d", project_id)
	resp, err := client.requestDoWithAuth("GET", url)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal server error",
		})
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal server error",
		})
	}

	if resp.StatusCode != 200 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"data":    "Project does not exist",
		})
	}

	var result ProjectBasicDetails
	err = json.Unmarshal(data, &result)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal server error",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    &result,
	})
}

func GetLanguageDetails(c *fiber.Ctx) error {
	project_id, err := c.ParamsInt("id")
	var token = os.Getenv("GITLAB_TOKEN")

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Please pass valid project id to get the data",
		})
	}

	var client = NewClient(token)
	url := fmt.Sprintf(GitLabAPI+"/%d/languages", project_id)
	resp, err := client.requestDoWithAuth("GET", url)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal Server Error",
		})
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal Server Error",
		})
	}

	var r ProjectLanguage

	var result map[string]float64
	err = json.Unmarshal(data, &result)

	r.Language = result

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal Server Error",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    &r,
	})
}

func GetMemberDetails(c *fiber.Ctx) error {
	project_id, err := c.ParamsInt("id")
	var token = os.Getenv("GITLAB_TOKEN")

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Please pass valid project id to get the data",
		})
	}

	var client = NewClient(token)
	url := fmt.Sprintf(GitLabAPI+"/%d/members", project_id)
	resp, err := client.requestDoWithAuth("GET", url)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal Server Error",
		})
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal Server Error",
		})
	}

	var result []*Members
	err = json.Unmarshal(data, &result)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"data":    "Internal Server Error",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    result,
	})
}