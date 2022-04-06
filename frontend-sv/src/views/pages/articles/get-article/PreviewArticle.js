import React from 'react'
import {
  CCard,
  CCardBody,
  CCardTitle,
  CCardHeader,
  CCardSubtitle,
  CCardText,
  CRow,
  CContainer,
} from '@coreui/react'
import { useState, useEffect } from 'react'
import ReactPaginate from 'react-paginate'
import './PreviewArticle.css'
import axios from '../../../../API'

function PreviewArticle() {
  const [articles, setArticles] = useState([])
  const [draftArticles, setDraft] = useState([])
  const [page, setPage] = useState(0)

  const perPage = 3
  const pageVisited = page * perPage

  const displayArticle = draftArticles
    .slice(pageVisited, pageVisited + perPage)
    .map((data, index) => {
      return (
        <CRow className="mt-4" key={index}>
          <CCard>
            <CCardHeader className="text-center text-white" style={{ backgroundColor: '#007bff' }}>
              <strong>Article #{data.id}</strong>
            </CCardHeader>
            <CCardBody>
              <CCardTitle>Title: {data.title}</CCardTitle>
              <CCardSubtitle>Category: {data.category}</CCardSubtitle>
              <CCardText>{data.content}</CCardText>
            </CCardBody>
          </CCard>
        </CRow>
      )
    })
  const getArticles = async () => {
    try {
      await axios.get(`/articles/`).then((res) => {
        setArticles(res.data.data)
      })
    } catch (error) {
      console.log(error)
    }
  }

  const pageCount = Math.ceil(draftArticles.length / perPage)

  const changePage = ({ selected }) => {
    setPage(selected)
  }

  const checkArticleStatus = async () => {
    var draftArticle = await articles.filter((e) => {
      return e.status === 'Publish'
    })
    setDraft(draftArticle)
  }

  useEffect(() => {
    getArticles()
  }, [])

  useEffect(() => {
    checkArticleStatus()
  }, [articles])

  return (
    <>
      {displayArticle}

      <CContainer className="mt-5">
        <ReactPaginate
          previousLabel={'Previous'}
          nextLabel={'Next'}
          pageCount={pageCount}
          onPageChange={changePage}
          containerClassName={'paginationButton'}
          previousLinkClassName={'previousButton'}
          nextLinkClassName={'nextButton'}
          activeClassName={'activeButton'}
        />
      </CContainer>
    </>
  )
}

export default PreviewArticle
