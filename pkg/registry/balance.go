package registry

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"github.com/wecab/go-epp-fi/pkg/epp"
)

func (s *Client) Balance() (int, error) {
	reqID := createRequestID(reqIDLength)
	balanceReq := epp.APIBalance{}
	balanceReq.Xmlns = epp.EPPNamespace
	balanceReq.Command.ClTRID = reqID

	balanceData, err := xml.Marshal(balanceReq)
	if err != nil {
		return -1, err
	}

	balanceRawResp, err := s.Send(balanceData)
	if err != nil {
		s.logAPIConnectionError(err, "requestID", reqID)
		return -1, err
	}

	var balanceResult epp.APIResult
	if err = xml.Unmarshal(balanceRawResp, &balanceResult); err != nil {
		return -1, err
	}

	if balanceResult.Response.Result.Code != 1000 {
		return -1, errors.New("Request failed: " + balanceResult.Response.Result.Msg)
	}

	return balanceResult.Response.ResData.BalanceAmount, nil
}
