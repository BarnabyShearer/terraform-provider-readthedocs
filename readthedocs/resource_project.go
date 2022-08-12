package readthedocs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	rtd "github.com/BarnabyShearer/readthedocs/v3"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Description:   "A readthedocs.org project.",
		CreateContext: resourceProjectCreate,
		UpdateContext: resourceProjectUpdate,
		ReadContext:   resourceProjectRead,
		DeleteContext: resourceProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The Slug of the project.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Project name.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL of repository to clone documentation from.",
			},
			"programming_language": {
				Type:        schema.TypeString,
				Default:     "py",
				Optional:    true,
				Description: "Given like `py`, `js` etc.",
			},
			"language": {
				Type:        schema.TypeString,
				Default:     "en",
				Optional:    true,
				Description: "Given like `en`, `fr` etc.",
			},
			"default_version": {
				Type:        schema.TypeString,
				Default:     "latest",
				Optional:    true,
				Description: "Version of documentation to show by default.",
			},
			"default_branch": {
				Type:        schema.TypeString,
				Default:     "main",
				Optional:    true,
				Description: "Branch to build from.",
			},
			"analytics_code": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Google Analytics code for tracking views.",
			},
			"analytics_disabled": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Disable Google Analytics.",
			},
			"show_version_warning": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				Description: "Warn when viewing old versions.",
			},
			"single_version": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Only show single version.",
			},
			"external_builds_enabled": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Build PRs.",
			},
			"organization": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "ReadTheDocs for Business organization where the project should be created. Only valid when using Read The Docs for Business.",
			},
			"teams": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				Description: "Team slugs the project will belong to. Only valid when using Read The Docs for Business.",
			},
		},
	}
}

func updateReqest(d *schema.ResourceData) rtd.CreateUpdateProject {
	return rtd.CreateUpdateProject{
		CreateProject: rtd.CreateProject{
			Name: d.Get("name").(string),
			Repository: rtd.Repository{
				URL:  d.Get("repository").(string),
				Type: "git",
			},
			Homepage:            "",
			ProgrammingLanguage: d.Get("programming_language").(string),
			Language:            d.Get("language").(string),
			Organization:        d.Get("organization").(string),
			Teams:               d.Get("teams").(string),
		},
		DefaultVersion:        d.Get("default_version").(string),
		DefaultBranch:         d.Get("default_branch").(string),
		AnalyticsCode:         d.Get("analytics_code").(string),
		AnalyticsDisabled:     d.Get("analytics_disabled").(bool),
		ShowVersionWarning:    d.Get("show_version_warning").(bool),
		SingleVersion:         d.Get("single_version").(bool),
		ExternalBuildsEnabled: d.Get("external_builds_enabled").(bool),
	}
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*rtd.Client)
	project, err := client.GetProject(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", project.Name)
	d.Set("repository", project.Repository.URL)
	d.Set("programming_language", project.ProgrammingLanguage.Code)
	d.Set("language", project.Language.Code)
	d.Set("default_version", project.DefaultVersion)
	d.Set("default_branch", project.DefaultBranch)
	d.SetId(project.Slug)
	return nil
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*rtd.Client)
	slug, err := client.CreateProject(ctx, updateReqest(d))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(slug)
	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*rtd.Client)
	err := client.UpdateProject(ctx, d.Id(), updateReqest(d))
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*rtd.Client)
	err := client.DeleteProject(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
