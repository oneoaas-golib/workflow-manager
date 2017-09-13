## workflow-manager/cloudformation

Cloudformation template for setting up AWS resources required by workflow-manager.

### Description

In order for workflow manager to get updates about running jobs, it needs to see events in the ECS cluster associated with the Batch environment.
The cloudformation template in this folder sets up a Cloudwatch Events rule that forwards ECS cluster events to an SQS queue.
This SQS queue is then monitored by workflow-manager.

``` bash
aws --region us-east-1 cloudformation create-stack \
  --stack-name workflow-manager \
  --template-body file://template.yml \
  --parameters \
    ParameterKey=ECSClusterARN,ParameterValue=arn:aws:ecs:us-east-1:589690932525:cluster/batch-east-dev_Batch_9c4a97f9-d9ab-3928-9c38-9e32c66a116c

aws --region us-east-1 cloudformation update-stack \
  --stack-name workflow-manager \
  --template-body file://template.yml \
  --parameters \
    ParameterKey=ECSClusterARN,ParameterValue=arn:aws:ecs:us-east-1:589690932525:cluster/batch-east-dev_Batch_9c4a97f9-d9ab-3928-9c38-9e32c66a116c
```
