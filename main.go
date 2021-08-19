package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	//get the List of available buckets
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//Iterate through each bucket and get the list of available objects
	for _, bucket := range buckets {
		fmt.Println("---->", bucket.Name)

		//List the objects in each bucket
		objList := minioClient.ListObjects(context.Background(), bucket.Name, minio.ListObjectsOptions{})
		for obj := range objList {
			fmt.Println(obj.Key)
		}
	}

	//get the object
	bucketName := "durga"
	objectKey := "2021-8-13.json"
	minObj, err := minioClient.GetObject(context.Background(), bucketName, objectKey, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println("Error in getting the object:", err)
		return
	}
	bytes, err := ioutil.ReadAll(minObj)
	if err != nil {
		fmt.Println("Error in reading the content from minObj:", err)
		return
	}

	report := []Report{}
	err = json.Unmarshal(bytes, &report)
	if err != nil {
		fmt.Println("Error in converting the content to json:", err)
		return
	}
	fmt.Println(len(report))
	fmt.Println(report[0].Version, report[0].ReportUuid)

	//Put the object
	bytes, err = json.Marshal(report[0])
	if err != nil {
		fmt.Println("Error in converting the content to bytes:", err)
		return
	}
	sReader := strings.NewReader(string(bytes))
	objName := fmt.Sprintf("%s_%s.json", "test", time.Now().Format("2006_01_02_15_04_05.999999"))
	info, err := minioClient.PutObject(context.Background(), bucketName, objName, sReader, -1, minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Error in putting the object:", err)
		return
	}
	fmt.Println("Success", info.Size)

}

type Report struct {
	Version     string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	OtherChecks []string `protobuf:"bytes,19,rep,name=other_checks,json=otherChecks,proto3" json:"other_checks,omitempty" toml:"other_checks,omitempty" mapstructure:"other_checks,omitempty"`
	ReportUuid  string   `protobuf:"bytes,20,opt,name=report_uuid,json=reportUuid,proto3" json:"report_uuid,omitempty" toml:"report_uuid,omitempty" mapstructure:"report_uuid,omitempty"`
	NodeUuid    string   `protobuf:"bytes,21,opt,name=node_uuid,json=nodeUuid,proto3" json:"node_uuid,omitempty" toml:"node_uuid,omitempty" mapstructure:"node_uuid,omitempty"`
	JobUuid     string   `protobuf:"bytes,22,opt,name=job_uuid,json=jobUuid,proto3" json:"job_uuid,omitempty" toml:"job_uuid,omitempty" mapstructure:"job_uuid,omitempty"`
	NodeName    string   `protobuf:"bytes,23,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty" toml:"node_name,omitempty" mapstructure:"node_name,omitempty"`
	Environment string   `protobuf:"bytes,24,opt,name=environment,proto3" json:"environment,omitempty" toml:"environment,omitempty" mapstructure:"environment,omitempty"`
	Roles       []string `protobuf:"bytes,25,rep,name=roles,proto3" json:"roles,omitempty" toml:"roles,omitempty" mapstructure:"roles,omitempty"`
	Recipes     []string `protobuf:"bytes,26,rep,name=recipes,proto3" json:"recipes,omitempty" toml:"recipes,omitempty" mapstructure:"recipes,omitempty"`
	//EndTime             string   `protobuf:"bytes,27,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" toml:"end_time,omitempty" mapstructure:"end_time,omitempty"`
	Type                string   `protobuf:"bytes,28,opt,name=type,proto3" json:"type,omitempty" toml:"type,omitempty" mapstructure:"type,omitempty"`
	SourceId            string   `protobuf:"bytes,29,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty" toml:"source_id,omitempty" mapstructure:"source_id,omitempty"`
	SourceRegion        string   `protobuf:"bytes,30,opt,name=source_region,json=sourceRegion,proto3" json:"source_region,omitempty" toml:"source_region,omitempty" mapstructure:"source_region,omitempty"`
	SourceAccountId     string   `protobuf:"bytes,31,opt,name=source_account_id,json=sourceAccountId,proto3" json:"source_account_id,omitempty" toml:"source_account_id,omitempty" mapstructure:"source_account_id,omitempty"`
	PolicyName          string   `protobuf:"bytes,32,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty" toml:"policy_name,omitempty" mapstructure:"policy_name,omitempty"`
	PolicyGroup         string   `protobuf:"bytes,33,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty" toml:"policy_group,omitempty" mapstructure:"policy_group,omitempty"`
	OrganizationName    string   `protobuf:"bytes,34,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty" toml:"organization_name,omitempty" mapstructure:"organization_name,omitempty"`
	SourceFqdn          string   `protobuf:"bytes,35,opt,name=source_fqdn,json=sourceFqdn,proto3" json:"source_fqdn,omitempty" toml:"source_fqdn,omitempty" mapstructure:"source_fqdn,omitempty"`
	ChefTags            []string `protobuf:"bytes,36,rep,name=chef_tags,json=chefTags,proto3" json:"chef_tags,omitempty" toml:"chef_tags,omitempty" mapstructure:"chef_tags,omitempty"`
	Ipaddress           string   `protobuf:"bytes,37,opt,name=ipaddress,proto3" json:"ipaddress,omitempty" toml:"ipaddress,omitempty" mapstructure:"ipaddress,omitempty"`
	Fqdn                string   `protobuf:"bytes,38,opt,name=fqdn,proto3" json:"fqdn,omitempty" toml:"fqdn,omitempty" mapstructure:"fqdn,omitempty"`
	AutomateManagerId   string   `protobuf:"bytes,40,opt,name=automate_manager_id,json=automateManagerId,proto3" json:"automate_manager_id,omitempty" toml:"automate_manager_id,omitempty" mapstructure:"automate_manager_id,omitempty"`
	RunTimeLimit        float32  `protobuf:"fixed32,41,opt,name=run_time_limit,json=runTimeLimit,proto3" json:"run_time_limit,omitempty" toml:"run_time_limit,omitempty" mapstructure:"run_time_limit,omitempty"`
	AutomateManagerType string   `protobuf:"bytes,42,opt,name=automate_manager_type,json=automateManagerType,proto3" json:"automate_manager_type,omitempty" toml:"automate_manager_type,omitempty" mapstructure:"automate_manager_type,omitempty"`
	Status              string   `protobuf:"bytes,43,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty" mapstructure:"status,omitempty"`
	StatusMessage       string   `protobuf:"bytes,44,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty" toml:"status_message,omitempty" mapstructure:"status_message,omitempty"`
}
