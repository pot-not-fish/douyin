// Code generated by thriftgo (0.3.1). DO NOT EDIT.

package favorite_rpc

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type IsFavoriteReq struct {
	UserId  []int64 `thrift:"user_id,1" frugal:"1,default,list<i64>" json:"user_id"`
	VideoId []int64 `thrift:"video_id,2" frugal:"2,default,list<i64>" json:"video_id"`
}

func NewIsFavoriteReq() *IsFavoriteReq {
	return &IsFavoriteReq{}
}

func (p *IsFavoriteReq) InitDefault() {
	*p = IsFavoriteReq{}
}

func (p *IsFavoriteReq) GetUserId() (v []int64) {
	return p.UserId
}

func (p *IsFavoriteReq) GetVideoId() (v []int64) {
	return p.VideoId
}
func (p *IsFavoriteReq) SetUserId(val []int64) {
	p.UserId = val
}
func (p *IsFavoriteReq) SetVideoId(val []int64) {
	p.VideoId = val
}

var fieldIDToName_IsFavoriteReq = map[int16]string{
	1: "user_id",
	2: "video_id",
}

func (p *IsFavoriteReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_IsFavoriteReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *IsFavoriteReq) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.UserId = make([]int64, 0, size)
	for i := 0; i < size; i++ {
		var _elem int64
		if v, err := iprot.ReadI64(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.UserId = append(p.UserId, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *IsFavoriteReq) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.VideoId = make([]int64, 0, size)
	for i := 0; i < size; i++ {
		var _elem int64
		if v, err := iprot.ReadI64(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.VideoId = append(p.VideoId, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *IsFavoriteReq) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("IsFavoriteReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *IsFavoriteReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_id", thrift.LIST, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.I64, len(p.UserId)); err != nil {
		return err
	}
	for _, v := range p.UserId {
		if err := oprot.WriteI64(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *IsFavoriteReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("video_id", thrift.LIST, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.I64, len(p.VideoId)); err != nil {
		return err
	}
	for _, v := range p.VideoId {
		if err := oprot.WriteI64(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *IsFavoriteReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IsFavoriteReq(%+v)", *p)
}

func (p *IsFavoriteReq) DeepEqual(ano *IsFavoriteReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.UserId) {
		return false
	}
	if !p.Field2DeepEqual(ano.VideoId) {
		return false
	}
	return true
}

func (p *IsFavoriteReq) Field1DeepEqual(src []int64) bool {

	if len(p.UserId) != len(src) {
		return false
	}
	for i, v := range p.UserId {
		_src := src[i]
		if v != _src {
			return false
		}
	}
	return true
}
func (p *IsFavoriteReq) Field2DeepEqual(src []int64) bool {

	if len(p.VideoId) != len(src) {
		return false
	}
	for i, v := range p.VideoId {
		_src := src[i]
		if v != _src {
			return false
		}
	}
	return true
}

type IsFavoriteResp struct {
	StatusCode int16  `thrift:"status_code,1" frugal:"1,default,i16" json:"status_code"`
	StatusMsg  string `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
	IsFavorite []bool `thrift:"is_favorite,3" frugal:"3,default,list<bool>" json:"is_favorite"`
}

func NewIsFavoriteResp() *IsFavoriteResp {
	return &IsFavoriteResp{}
}

func (p *IsFavoriteResp) InitDefault() {
	*p = IsFavoriteResp{}
}

func (p *IsFavoriteResp) GetStatusCode() (v int16) {
	return p.StatusCode
}

func (p *IsFavoriteResp) GetStatusMsg() (v string) {
	return p.StatusMsg
}

func (p *IsFavoriteResp) GetIsFavorite() (v []bool) {
	return p.IsFavorite
}
func (p *IsFavoriteResp) SetStatusCode(val int16) {
	p.StatusCode = val
}
func (p *IsFavoriteResp) SetStatusMsg(val string) {
	p.StatusMsg = val
}
func (p *IsFavoriteResp) SetIsFavorite(val []bool) {
	p.IsFavorite = val
}

var fieldIDToName_IsFavoriteResp = map[int16]string{
	1: "status_code",
	2: "status_msg",
	3: "is_favorite",
}

func (p *IsFavoriteResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I16 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_IsFavoriteResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *IsFavoriteResp) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return err
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *IsFavoriteResp) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.StatusMsg = v
	}
	return nil
}

func (p *IsFavoriteResp) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.IsFavorite = make([]bool, 0, size)
	for i := 0; i < size; i++ {
		var _elem bool
		if v, err := iprot.ReadBool(); err != nil {
			return err
		} else {
			_elem = v
		}

		p.IsFavorite = append(p.IsFavorite, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *IsFavoriteResp) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("IsFavoriteResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *IsFavoriteResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status_code", thrift.I16, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI16(p.StatusCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *IsFavoriteResp) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status_msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.StatusMsg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *IsFavoriteResp) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("is_favorite", thrift.LIST, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.BOOL, len(p.IsFavorite)); err != nil {
		return err
	}
	for _, v := range p.IsFavorite {
		if err := oprot.WriteBool(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *IsFavoriteResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IsFavoriteResp(%+v)", *p)
}

func (p *IsFavoriteResp) DeepEqual(ano *IsFavoriteResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.StatusCode) {
		return false
	}
	if !p.Field2DeepEqual(ano.StatusMsg) {
		return false
	}
	if !p.Field3DeepEqual(ano.IsFavorite) {
		return false
	}
	return true
}

func (p *IsFavoriteResp) Field1DeepEqual(src int16) bool {

	if p.StatusCode != src {
		return false
	}
	return true
}
func (p *IsFavoriteResp) Field2DeepEqual(src string) bool {

	if strings.Compare(p.StatusMsg, src) != 0 {
		return false
	}
	return true
}
func (p *IsFavoriteResp) Field3DeepEqual(src []bool) bool {

	if len(p.IsFavorite) != len(src) {
		return false
	}
	for i, v := range p.IsFavorite {
		_src := src[i]
		if v != _src {
			return false
		}
	}
	return true
}

type FavoriteService interface {
	IsFavorite(ctx context.Context, request *IsFavoriteReq) (r *IsFavoriteResp, err error)
}

type FavoriteServiceClient struct {
	c thrift.TClient
}

func NewFavoriteServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FavoriteServiceClient {
	return &FavoriteServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewFavoriteServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FavoriteServiceClient {
	return &FavoriteServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewFavoriteServiceClient(c thrift.TClient) *FavoriteServiceClient {
	return &FavoriteServiceClient{
		c: c,
	}
}

func (p *FavoriteServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *FavoriteServiceClient) IsFavorite(ctx context.Context, request *IsFavoriteReq) (r *IsFavoriteResp, err error) {
	var _args FavoriteServiceIsFavoriteArgs
	_args.Request = request
	var _result FavoriteServiceIsFavoriteResult
	if err = p.Client_().Call(ctx, "IsFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type FavoriteServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      FavoriteService
}

func (p *FavoriteServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *FavoriteServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *FavoriteServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewFavoriteServiceProcessor(handler FavoriteService) *FavoriteServiceProcessor {
	self := &FavoriteServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("IsFavorite", &favoriteServiceProcessorIsFavorite{handler: handler})
	return self
}
func (p *FavoriteServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type favoriteServiceProcessorIsFavorite struct {
	handler FavoriteService
}

func (p *favoriteServiceProcessorIsFavorite) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FavoriteServiceIsFavoriteArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("IsFavorite", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := FavoriteServiceIsFavoriteResult{}
	var retval *IsFavoriteResp
	if retval, err2 = p.handler.IsFavorite(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing IsFavorite: "+err2.Error())
		oprot.WriteMessageBegin("IsFavorite", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("IsFavorite", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type FavoriteServiceIsFavoriteArgs struct {
	Request *IsFavoriteReq `thrift:"request,1" frugal:"1,default,IsFavoriteReq" json:"request"`
}

func NewFavoriteServiceIsFavoriteArgs() *FavoriteServiceIsFavoriteArgs {
	return &FavoriteServiceIsFavoriteArgs{}
}

func (p *FavoriteServiceIsFavoriteArgs) InitDefault() {
	*p = FavoriteServiceIsFavoriteArgs{}
}

var FavoriteServiceIsFavoriteArgs_Request_DEFAULT *IsFavoriteReq

func (p *FavoriteServiceIsFavoriteArgs) GetRequest() (v *IsFavoriteReq) {
	if !p.IsSetRequest() {
		return FavoriteServiceIsFavoriteArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *FavoriteServiceIsFavoriteArgs) SetRequest(val *IsFavoriteReq) {
	p.Request = val
}

var fieldIDToName_FavoriteServiceIsFavoriteArgs = map[int16]string{
	1: "request",
}

func (p *FavoriteServiceIsFavoriteArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *FavoriteServiceIsFavoriteArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceIsFavoriteArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteServiceIsFavoriteArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewIsFavoriteReq()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FavoriteServiceIsFavoriteArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("IsFavorite_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FavoriteServiceIsFavoriteArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FavoriteServiceIsFavoriteArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FavoriteServiceIsFavoriteArgs(%+v)", *p)
}

func (p *FavoriteServiceIsFavoriteArgs) DeepEqual(ano *FavoriteServiceIsFavoriteArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *FavoriteServiceIsFavoriteArgs) Field1DeepEqual(src *IsFavoriteReq) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type FavoriteServiceIsFavoriteResult struct {
	Success *IsFavoriteResp `thrift:"success,0,optional" frugal:"0,optional,IsFavoriteResp" json:"success,omitempty"`
}

func NewFavoriteServiceIsFavoriteResult() *FavoriteServiceIsFavoriteResult {
	return &FavoriteServiceIsFavoriteResult{}
}

func (p *FavoriteServiceIsFavoriteResult) InitDefault() {
	*p = FavoriteServiceIsFavoriteResult{}
}

var FavoriteServiceIsFavoriteResult_Success_DEFAULT *IsFavoriteResp

func (p *FavoriteServiceIsFavoriteResult) GetSuccess() (v *IsFavoriteResp) {
	if !p.IsSetSuccess() {
		return FavoriteServiceIsFavoriteResult_Success_DEFAULT
	}
	return p.Success
}
func (p *FavoriteServiceIsFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*IsFavoriteResp)
}

var fieldIDToName_FavoriteServiceIsFavoriteResult = map[int16]string{
	0: "success",
}

func (p *FavoriteServiceIsFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteServiceIsFavoriteResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceIsFavoriteResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteServiceIsFavoriteResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewIsFavoriteResp()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FavoriteServiceIsFavoriteResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("IsFavorite_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FavoriteServiceIsFavoriteResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *FavoriteServiceIsFavoriteResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FavoriteServiceIsFavoriteResult(%+v)", *p)
}

func (p *FavoriteServiceIsFavoriteResult) DeepEqual(ano *FavoriteServiceIsFavoriteResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *FavoriteServiceIsFavoriteResult) Field0DeepEqual(src *IsFavoriteResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
