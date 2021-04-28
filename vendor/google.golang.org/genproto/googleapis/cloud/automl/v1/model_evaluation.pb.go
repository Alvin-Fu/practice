// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/model_evaluation.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Evaluation results of a model.
type ModelEvaluation struct {
	// Output only. Problem type specific evaluation metrics.
	//
	// Types that are valid to be assigned to Metrics:
	//	*ModelEvaluation_ClassificationEvaluationMetrics
	//	*ModelEvaluation_TranslationEvaluationMetrics
	//	*ModelEvaluation_ImageObjectDetectionEvaluationMetrics
	//	*ModelEvaluation_TextSentimentEvaluationMetrics
	//	*ModelEvaluation_TextExtractionEvaluationMetrics
	Metrics isModelEvaluation_Metrics `protobuf_oneof:"metrics"`
	// Output only. Resource name of the model evaluation.
	// Format:
	//
	// `projects/{project_id}/locations/{location_id}/models/{model_id}/modelEvaluations/{model_evaluation_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The ID of the annotation spec that the model evaluation
	// applies to. The The ID is empty for the overall model evaluation.
	AnnotationSpecId string `protobuf:"bytes,2,opt,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only. The value of
	// [display_name][google.cloud.automl.v1.AnnotationSpec.display_name]
	// at the moment when the model was trained. Because this field returns a
	// value at model training time, for different models trained from the same
	// dataset, the values may differ, since display names could had been changed
	// between the two model's trainings.
	DisplayName string `protobuf:"bytes,15,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. Timestamp when this model evaluation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The number of examples used for model evaluation, i.e. for
	// which ground truth from time of model creation is compared against the
	// predicted annotations created by the model.
	// For overall ModelEvaluation (i.e. with annotation_spec_id not set) this is
	// the total number of all examples used for evaluation.
	// Otherwise, this is the count of examples that according to the ground
	// truth were annotated by the
	//
	// [annotation_spec_id][google.cloud.automl.v1.ModelEvaluation.annotation_spec_id].
	EvaluatedExampleCount int32    `protobuf:"varint,6,opt,name=evaluated_example_count,json=evaluatedExampleCount,proto3" json:"evaluated_example_count,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ModelEvaluation) Reset()         { *m = ModelEvaluation{} }
func (m *ModelEvaluation) String() string { return proto.CompactTextString(m) }
func (*ModelEvaluation) ProtoMessage()    {}
func (*ModelEvaluation) Descriptor() ([]byte, []int) {
	return fileDescriptor_008481175b84a2ca, []int{0}
}

func (m *ModelEvaluation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelEvaluation.Unmarshal(m, b)
}
func (m *ModelEvaluation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelEvaluation.Marshal(b, m, deterministic)
}
func (m *ModelEvaluation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelEvaluation.Merge(m, src)
}
func (m *ModelEvaluation) XXX_Size() int {
	return xxx_messageInfo_ModelEvaluation.Size(m)
}
func (m *ModelEvaluation) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelEvaluation.DiscardUnknown(m)
}

var xxx_messageInfo_ModelEvaluation proto.InternalMessageInfo

type isModelEvaluation_Metrics interface {
	isModelEvaluation_Metrics()
}

type ModelEvaluation_ClassificationEvaluationMetrics struct {
	ClassificationEvaluationMetrics *ClassificationEvaluationMetrics `protobuf:"bytes,8,opt,name=classification_evaluation_metrics,json=classificationEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_TranslationEvaluationMetrics struct {
	TranslationEvaluationMetrics *TranslationEvaluationMetrics `protobuf:"bytes,9,opt,name=translation_evaluation_metrics,json=translationEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_ImageObjectDetectionEvaluationMetrics struct {
	ImageObjectDetectionEvaluationMetrics *ImageObjectDetectionEvaluationMetrics `protobuf:"bytes,12,opt,name=image_object_detection_evaluation_metrics,json=imageObjectDetectionEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_TextSentimentEvaluationMetrics struct {
	TextSentimentEvaluationMetrics *TextSentimentEvaluationMetrics `protobuf:"bytes,11,opt,name=text_sentiment_evaluation_metrics,json=textSentimentEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_TextExtractionEvaluationMetrics struct {
	TextExtractionEvaluationMetrics *TextExtractionEvaluationMetrics `protobuf:"bytes,13,opt,name=text_extraction_evaluation_metrics,json=textExtractionEvaluationMetrics,proto3,oneof"`
}

func (*ModelEvaluation_ClassificationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_TranslationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_ImageObjectDetectionEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_TextSentimentEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_TextExtractionEvaluationMetrics) isModelEvaluation_Metrics() {}

func (m *ModelEvaluation) GetMetrics() isModelEvaluation_Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ModelEvaluation) GetClassificationEvaluationMetrics() *ClassificationEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_ClassificationEvaluationMetrics); ok {
		return x.ClassificationEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetTranslationEvaluationMetrics() *TranslationEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_TranslationEvaluationMetrics); ok {
		return x.TranslationEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetImageObjectDetectionEvaluationMetrics() *ImageObjectDetectionEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_ImageObjectDetectionEvaluationMetrics); ok {
		return x.ImageObjectDetectionEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetTextSentimentEvaluationMetrics() *TextSentimentEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_TextSentimentEvaluationMetrics); ok {
		return x.TextSentimentEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetTextExtractionEvaluationMetrics() *TextExtractionEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_TextExtractionEvaluationMetrics); ok {
		return x.TextExtractionEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelEvaluation) GetAnnotationSpecId() string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return ""
}

func (m *ModelEvaluation) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ModelEvaluation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ModelEvaluation) GetEvaluatedExampleCount() int32 {
	if m != nil {
		return m.EvaluatedExampleCount
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ModelEvaluation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ModelEvaluation_ClassificationEvaluationMetrics)(nil),
		(*ModelEvaluation_TranslationEvaluationMetrics)(nil),
		(*ModelEvaluation_ImageObjectDetectionEvaluationMetrics)(nil),
		(*ModelEvaluation_TextSentimentEvaluationMetrics)(nil),
		(*ModelEvaluation_TextExtractionEvaluationMetrics)(nil),
	}
}

func init() {
	proto.RegisterType((*ModelEvaluation)(nil), "google.cloud.automl.v1.ModelEvaluation")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/model_evaluation.proto", fileDescriptor_008481175b84a2ca)
}

var fileDescriptor_008481175b84a2ca = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x61, 0x6b, 0xd3, 0x40,
	0x18, 0xc7, 0xcd, 0x70, 0x9b, 0xbb, 0x4e, 0x26, 0x07, 0xce, 0xae, 0x8c, 0x6d, 0x1d, 0x28, 0x95,
	0xcd, 0x84, 0xaa, 0x4c, 0xc8, 0xf4, 0x85, 0xab, 0x45, 0x07, 0x56, 0xa5, 0x1b, 0x7d, 0x21, 0x85,
	0x70, 0xbd, 0x3c, 0x0b, 0x91, 0x4b, 0x2e, 0xe4, 0x2e, 0xa5, 0x82, 0xef, 0x44, 0xbf, 0x83, 0x9f,
	0x40, 0xd8, 0x47, 0xf1, 0xa3, 0xf8, 0x29, 0x24, 0x77, 0x49, 0x9b, 0x8e, 0x24, 0xf8, 0x2e, 0xc9,
	0xff, 0xf7, 0xfc, 0xef, 0x7f, 0xcf, 0x3d, 0x39, 0xf4, 0xc4, 0xe3, 0xdc, 0x63, 0x60, 0x51, 0xc6,
	0x13, 0xd7, 0x22, 0x89, 0xe4, 0x01, 0xb3, 0xa6, 0x5d, 0x2b, 0xe0, 0x2e, 0x30, 0x07, 0xa6, 0x84,
	0x25, 0x44, 0xfa, 0x3c, 0x34, 0xa3, 0x98, 0x4b, 0x8e, 0xb7, 0x35, 0x6e, 0x2a, 0xdc, 0xd4, 0xb8,
	0x39, 0xed, 0xb6, 0x76, 0x33, 0x1b, 0x12, 0xf9, 0x16, 0x09, 0x43, 0x2e, 0x55, 0x91, 0xd0, 0x55,
	0xad, 0x9d, 0x82, 0x1a, 0x83, 0xe0, 0x49, 0x4c, 0x21, 0x93, 0x8e, 0x2a, 0xd6, 0xa7, 0x8c, 0x08,
	0xe1, 0x5f, 0xf9, 0xb4, 0xb0, 0x7a, 0xeb, 0x51, 0x05, 0xec, 0x82, 0x04, 0x5a, 0xe0, 0x8e, 0x2b,
	0x38, 0x09, 0x33, 0xe9, 0xc0, 0x4c, 0xc6, 0xa4, 0x48, 0x1f, 0xd5, 0xd1, 0x02, 0x42, 0xe9, 0x07,
	0x10, 0xca, 0x0c, 0xee, 0x54, 0xc1, 0x31, 0x09, 0x05, 0x2b, 0x86, 0xdd, 0xcf, 0x48, 0xf5, 0x36,
	0x49, 0xae, 0xac, 0xd4, 0x47, 0x48, 0x12, 0x44, 0x1a, 0x38, 0xfc, 0xbd, 0x8e, 0xb6, 0x06, 0x69,
	0x9b, 0xfb, 0xf3, 0x2e, 0xe3, 0x1f, 0x06, 0x6a, 0x2f, 0x6f, 0xbd, 0x70, 0x06, 0x4e, 0x00, 0x32,
	0xf6, 0xa9, 0x68, 0xde, 0x39, 0x30, 0x3a, 0x8d, 0xa7, 0x2f, 0xcc, 0xf2, 0xc3, 0x30, 0x7b, 0x4b,
	0x06, 0x0b, 0xf7, 0x81, 0x2e, 0x7f, 0x77, 0x6b, 0xb8, 0x4f, 0xeb, 0x11, 0xfc, 0x0d, 0xed, 0x15,
	0x76, 0x54, 0x96, 0x61, 0x43, 0x65, 0x78, 0x5e, 0x95, 0xe1, 0x72, 0x51, 0x5d, 0x16, 0x60, 0x57,
	0xd6, 0xe8, 0xf8, 0x97, 0x81, 0x1e, 0xfb, 0x01, 0xf1, 0xc0, 0xe1, 0x93, 0x2f, 0x40, 0xa5, 0x33,
	0x3f, 0xe0, 0xb2, 0x24, 0x9b, 0x2a, 0xc9, 0xab, 0xaa, 0x24, 0xe7, 0xa9, 0xd1, 0x47, 0xe5, 0xf3,
	0x26, 0xb7, 0x29, 0x8b, 0xf4, 0xd0, 0xff, 0x1f, 0x10, 0x7f, 0x37, 0x50, 0x7b, 0x79, 0x32, 0xca,
	0x32, 0x35, 0x54, 0xa6, 0x93, 0xca, 0xee, 0xc0, 0x4c, 0x5e, 0xe4, 0xf5, 0x65, 0x61, 0xf6, 0x64,
	0x2d, 0x81, 0x7f, 0x1a, 0xe8, 0xf0, 0xc6, 0x34, 0x97, 0xc5, 0xb8, 0x5b, 0x3f, 0x28, 0x69, 0x8c,
	0xfe, 0xdc, 0xa0, 0x74, 0x50, 0x64, 0x3d, 0x82, 0x31, 0xba, 0x1d, 0x92, 0x00, 0x9a, 0xc6, 0x81,
	0xd1, 0xd9, 0x18, 0xaa, 0x67, 0x7c, 0x8c, 0xf0, 0xe2, 0x0e, 0x70, 0x44, 0x04, 0xd4, 0xf1, 0xdd,
	0xe6, 0x8a, 0x22, 0xee, 0x2d, 0x94, 0x8b, 0x08, 0xe8, 0xb9, 0x8b, 0xdb, 0x68, 0xd3, 0xf5, 0x45,
	0xc4, 0xc8, 0x57, 0x47, 0x39, 0x6d, 0x29, 0xae, 0x91, 0x7d, 0xfb, 0x90, 0x1a, 0x9e, 0xa2, 0x06,
	0x8d, 0x81, 0x48, 0x70, 0xd2, 0x7e, 0x34, 0x57, 0xd5, 0xae, 0x5a, 0xf9, 0xae, 0xf2, 0x1f, 0xcc,
	0xbc, 0xcc, 0x7f, 0xb0, 0x21, 0xd2, 0x78, 0xfa, 0x01, 0x9f, 0xa0, 0x07, 0x59, 0x67, 0xc0, 0x75,
	0x60, 0x46, 0x82, 0x88, 0x81, 0x43, 0x79, 0x12, 0xca, 0xe6, 0xda, 0x81, 0xd1, 0x59, 0x1d, 0xde,
	0x9f, 0xcb, 0x7d, 0xad, 0xf6, 0x52, 0xf1, 0x6c, 0x03, 0xad, 0x67, 0x6d, 0x3c, 0xbb, 0x36, 0x50,
	0x8b, 0xf2, 0xa0, 0xa2, 0x8d, 0x9f, 0x8c, 0xcf, 0x2f, 0x33, 0xc5, 0xe3, 0x8c, 0x84, 0x9e, 0xc9,
	0x63, 0xcf, 0xf2, 0x20, 0x54, 0xc1, 0x2c, 0x2d, 0x91, 0xc8, 0x17, 0x37, 0x2f, 0x8d, 0x53, 0xfd,
	0x74, 0xbd, 0xb2, 0xfd, 0x56, 0x97, 0xf7, 0x94, 0xf1, 0xeb, 0x44, 0xf2, 0xc1, 0x7b, 0x73, 0xd4,
	0xfd, 0x93, 0x0b, 0x63, 0x25, 0x8c, 0x95, 0xc0, 0xc6, 0xa3, 0xee, 0xdf, 0x95, 0x1d, 0x2d, 0xd8,
	0xb6, 0x52, 0x6c, 0x5b, 0xd7, 0xd8, 0xf6, 0xa8, 0x3b, 0x59, 0x53, 0xcb, 0x3e, 0xfb, 0x17, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0xad, 0xb6, 0x91, 0xda, 0x05, 0x00, 0x00,
}
